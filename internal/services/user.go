package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/SMTP"
	"GeneReport_platform/internal/contracts/sharingPlatformContract"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appContext"
	"GeneReport_platform/pkg/appErrors"
	"GeneReport_platform/pkg/auth"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	UserServ *userService
)

type userService struct {
	iUserBase
}

// 用户基础接口
type iUserBase interface {
	IsNewUser(userAddress string) bool
	EnsureUserExists(userAddress string) error
	GetNonce(address string) (string, error)
}

func RegisterUserService() {
	UserServ = &userService{}
}

// IsNewUser 判断用户是否为新用户
func (u *userService) IsNewUser(userAddress string) (bool, appErrors.IAppError) {
	isNew, err := dao.GetUserDao().IsExist(userAddress)
	if err != nil {
		log.Printf("数据库内部错误为%v\n", err)
		return false, appErrors.New(http.StatusInternalServerError, "数据库内部错误", err)
	}
	return !isNew, nil
}

// EnsureUserExists 确保用户存在
func (u *userService) EnsureUserExists(userAddress string) error {
	if isNew, err := UserServ.IsNewUser(userAddress); err != nil {
		return err
	} else if isNew {
		if err := dao.GetUserDao().CreateUser(userAddress); err != nil {
			return appErrors.New(http.StatusInternalServerError, "新建用户失败", err)
		}
	}
	return nil
}

func (u *userService) GetNonce(address string) (string, appErrors.IAppError) {
	strCmd := configs.RedisClient.GetEx(appContext.NewTimeoutContext(), address, 3*time.Minute)
	//若当前查询得到的nonce还未过期，则直接返回，并更新过期时间
	if strCmd.Err() == nil {
		defer log.Printf(" 用户地址: %v; nonce: %v\n", address, strCmd.Val())
		return strCmd.Val(), nil
	} else if appErrors.Is(strCmd.Err(), redis.Nil) {
		//已过期或不存在，执行生成nonce
		nonce := auth.GenerateNonce()

		if err := configs.RedisClient.SetEX(appContext.NewTimeoutContext(), address, nonce, 3*time.Minute).Err(); err == nil {
			defer log.Printf(" 用户地址是: %v; nonce: %v\n", address, nonce)
			return nonce, nil
		}
	}
	// Redis 查询出错或不可用
	defer log.Printf("查询 nonce 失败，用户地址: %v, 错误: %v", address, strCmd.Err())
	return "", appErrors.New(http.StatusServiceUnavailable, "获取nonce失败", strCmd.Err())
}

func (u *userService) UpdateUser(toUpdate dao.UpdateUser) appErrors.IAppError {
	if err := dao.GetUserDao().UpdateUser(toUpdate); err != nil {
		return appErrors.New(http.StatusServiceUnavailable, "服务器内部错误", err)
	}
	return nil
}

func (u *userService) GetUserInfoByID(address string) (dto.UserInfoResp, error) {
	user, err := dao.GetUserDao().GetUser(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserInfoResp{}, appErrors.New(http.StatusNotFound, "用户不存在", err)
		} else {
			return dto.UserInfoResp{}, appErrors.New(http.StatusInternalServerError, "服务器内部错误", err)
		}
	}
	// 映射转换dto
	var userInfo dto.UserInfoResp
	mapstructure.Decode(user, &userInfo)
	return userInfo, nil
}

func (u *userService) VerifyInstitutionEmail(institutionName, email string) (isValid bool, err error) {
	suffix, err := dao.GetUserDao().GetEmailSuffix(institutionName)
	if err != nil {
		return false, appErrors.New(http.StatusInternalServerError, "服务器内部错误", err)
	}
	if strings.Split(email, "@")[1] == suffix[1:] {
		return true, nil
	}
	return false, nil
}

// SendEmailCode 发送短信验证码
// Redis存储格式
// key:"email_code:{email}" value:{"code":{code},"institution":{institution}} expire: 3min
// key:"email_attempts:{email}" value:0 expire: 4min
func (u *userService) SendEmailCode(userAddress, institutionName, email string) error {
	//验证邮箱是否已被使用
	isUsed, err := dao.GetUserDao().IsEmailUsed(email)
	if err != nil {
		return appErrors.New(http.StatusInternalServerError, "服务器内部错误", err)
	}
	if isUsed {
		return appErrors.New(http.StatusConflict, "邮箱已被使用")
	}

	//验证邮箱后缀合法性
	suffix, err := dao.GetUserDao().GetEmailSuffix(institutionName)
	if err != nil {
		return appErrors.New(http.StatusInternalServerError, "服务器内部错误", err)
	}
	if strings.Split(email, "@")[1] != suffix[1:] {
		return appErrors.New(http.StatusBadRequest, "邮箱格式不正确")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	codeToSave := SMTP.GenerateVerifyCode()

	User, _ := dao.GetUserDao().GetUser(userAddress)

	// 发送验证码邮件
	err = SMTP.SendEmailCode(codeToSave, User.Name, userAddress, email)
	if err != nil {
		return appErrors.New(http.StatusInternalServerError, "SMTP服务繁忙", err)
	}
	// 创建 Pipeline
	pipe := configs.RedisClient.Pipeline()

	// 添加 HSet 命令
	setCmd := pipe.HSet(ctx, "email_code:"+email, map[string]interface{}{
		"code":        codeToSave,
		"institution": institutionName,
	})
	pipe.Expire(ctx, "email_code:"+email, 3*time.Minute)
	// 添加 Expire 命令
	expiryCmd := pipe.Set(ctx, "email_attempts:"+email, 0, 4*time.Minute)

	// 执行 Pipeline
	_, err = pipe.Exec(ctx)
	if err != nil || expiryCmd.Err() != nil || setCmd.Err() != nil {
		return appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}
	return nil
}

// VerifyEmailCode 验证短信验证码
// Redis存储格式
// key:"email_code:{email}" value:{"code":{code},"institution":{institution}} expire: 3min
// key:"email_attempts:{email}" value:0 expire: 4min
func (u *userService) VerifyEmailCode(email string, code string, userAddress string) (bool, error) {
	// 创建 Pipeline
	pipe := configs.RedisClient.Pipeline()

	getAllCmd := pipe.HGetAll(context.Background(), "email_code:"+email)
	incrCmd := pipe.IncrBy(context.Background(), "email_attempts:"+email, 1)

	// 执行 Pipeline
	_, err := pipe.Exec(context.Background())
	if err != nil || getAllCmd.Err() != nil || incrCmd.Err() != nil {
		return false, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}

	if vals, err := getAllCmd.Result(); err != nil {
		return false, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	} else if vals["code"] == "" {
		return false, appErrors.New(http.StatusBadRequest, "验证码已过期")
	} else {
		attempts, _ := incrCmd.Result()
		//若尝试次数过多
		if attempts >= 5 {
			return false, appErrors.New(http.StatusTooManyRequests, "验证码错误次数过多，请三分钟后再试")
		}
		//若验证码错误
		if vals["code"] != code {
			return false, appErrors.New(http.StatusBadRequest, "验证码错误", errors.New("验证码错误"))
		}

		//验证码正确，更新到用户表字段
		err = dao.GetUserDao().UpdateUserInstitution(userAddress, vals["institution"], email)
		if err != nil {
			return false, err
		}

		//验证码正确，调合约
		_, receipt, err := sharingPlatformContract.GetContractIns().SetUserAuthStatus(sharingPlatformContract.NewAdminTransactor(), common.HexToAddress(userAddress))
		if err != nil {
			return false, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
		}
		fmt.Printf("SetUserAuthStatus执行成功，交易hash为：%v\n", receipt.TransactionHash)
		return true, nil
	}

}
