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
	"strings"
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
func (u *userService) SendSMSCode(phone string) appErrors.IAppError {
	message := unisms.BuildMessage()
	message.SetTo(phone)
	message.SetSignature("林锐轩")
	message.SetTemplateId("pub_verif_en_ttl2")

	codeToSave := smsVerify.GenerateSMSCode()
	// 设置模板数据（code,ttl）
	message.SetTemplateData(map[string]string{"code": codeToSave, "ttl": "10"})
	// 发送短信
	_, err := smsVerify.UniSMSClient.Send(message)
	log.Println("是否被初始化:", smsVerify.IsInit)
	if err != nil {
		return appErrors.New(http.StatusServiceUnavailable, "服务繁忙，请稍后再试", err)
	}

	// 存redis
	if err = configs.RedisClient.SetEX(context.Background(), "SMS_phone:"+phone, codeToSave, time.Minute*10).Err(); err != nil {
		return appErrors.New(http.StatusServiceUnavailable, "内部错误", err)
	}
	return nil
}

// VerifySMSCode 验证短信验证码
func (u *userService) VerifySMSCode(phone string, code string) (bool, appErrors.IAppError) {
	//redis取验证码
	strCmd := configs.RedisClient.Get(context.Background(), "SMS_phone:"+phone)

	if strCmd.Err() == nil {
		if strings.Split(strCmd.Val(), ":")[1] == code {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		return false, appErrors.New(http.StatusNotFound, "未发送验证码或验证码已过期", errors.New("未发送验证码或验证码已过期"))
	}

}
