package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/smsVerify"
	"GeneReport_platform/pkg/appContext"
	"GeneReport_platform/pkg/appErrors"
	"GeneReport_platform/pkg/auth"
	"context"
	"errors"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var (
	UserService *userService
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
	UserService = &userService{}
}

// IsNewUser 判断用户是否为新用户
func (u *userService) IsNewUser(userAddress string) (bool, appErrors.IAppError) {
	isNew, err := dao.UserDao.IsExist(userAddress)
	if err != nil {
		log.Printf("数据库内部错误为%v\n", err)
		return false, appErrors.New(http.StatusInternalServerError, "数据库内部错误", err)
	}
	return !isNew, nil
}

// EnsureUserExists 确保用户存在
func (u *userService) EnsureUserExists(userAddress string) appErrors.IAppError {
	if isNew, err := UserService.IsNewUser(userAddress); err != nil {
		return err
	} else if isNew {
		if err := dao.UserDao.CreateUser(userAddress); err != nil {
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

func (u *userService) UpdateUser(toUpdate dto.UpdateUser) appErrors.IAppError {
	if err := dao.UserDao.UpdateUser(toUpdate); err != nil {
		return appErrors.New(http.StatusServiceUnavailable, "服务器内部错误", err)
	}
	return nil
}

func (u *userService) GetUserInfo(address string) (dto.UpdateUser, appErrors.IAppError) {
	userInfo, err := dao.UserDao.GetUserInfo(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UpdateUser{}, appErrors.New(http.StatusNotFound, "用户不存在", err)
		} else {
			return dto.UpdateUser{}, appErrors.New(http.StatusInternalServerError, "服务器内部错误", err)
		}
	}
	return userInfo, nil
}

// SendSMSCode 发送短信验证码
// Redis存储格式 key:"SMS_phone:{phone}" value:{"code":{code},"attempts":0} expire:10min
func (u *userService) SendSMSCode(phone string) error {
	message := unisms.BuildMessage()
	message.SetTo(phone)
	message.SetSignature("林锐轩")
	message.SetTemplateId("pub_verif_en_ttl2")

	codeToSave := smsVerify.GenerateSMSCode()
	// 设置模板数据（code,ttl）
	message.SetTemplateData(map[string]string{"code": codeToSave, "ttl": "10"})
	// 设置上下文超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 发送短信
	_, err := smsVerify.UniSMSClient.Send(message)

	if err != nil {
		return appErrors.New(http.StatusServiceUnavailable, "服务繁忙，请稍后再试", err)
	}

	// 创建 Pipeline
	pipe := configs.RedisClient.Pipeline()

	// 在 Pipeline 中添加 HSet 命令
	pipe.HSet(ctx, "SMS_phone:"+phone, map[string]interface{}{
		"code":     codeToSave,
		"attempts": 0,
	})

	// 在 Pipeline 中添加 Expire 命令
	pipe.Expire(ctx, "SMS_phone:"+phone, 10*time.Minute)

	// 执行 Pipeline
	_, err = pipe.Exec(ctx)
	if err != nil {
		return appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	}
	return nil
}

// VerifySMSCode 验证短信验证码
// Redis存储格式 key:"SMS_phone:{phone}" value:{"code":{code},"attempts":{attempts}} expire:10min
func (u *userService) VerifySMSCode(phone string, code string) (bool, error) {
	var correctCode string
	var attempts int64
	sliceCmd := configs.RedisClient.HMGet(context.Background(), "SMS_phone:"+phone, "code", "attempts")
	if err := sliceCmd.Err(); err != nil {
		return false, appErrors.New(http.StatusInternalServerError, "服务繁忙，请稍后再试", err)
	} else {
		correctCode = sliceCmd.Val()[0].(string)
		attempts = sliceCmd.Val()[1].(int64)
		if attempts >= 5 {
			return false, appErrors.New(http.StatusTooManyRequests, "验证码错误次数过多，请十分钟后再试", err)
		}

		if correctCode != code {
			defer configs.RedisClient.HIncrBy(context.Background(), "SMS_phone:"+phone, "attempts", 1)
			return false, appErrors.New(http.StatusBadRequest, "验证码错误", errors.New("验证码错误"))

		}
		return true, nil
	}
}
