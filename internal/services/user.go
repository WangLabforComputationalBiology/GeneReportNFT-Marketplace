package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appContext"
	"GeneReport_platform/pkg/auth"
	"GeneReport_platform/pkg/custom_errors"
	"errors"
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
func (u *userService) IsNewUser(userAddress string) (bool, custom_errors.IAppError) {
	isNew, err := dao.UserDao.IsExist(userAddress)
	if err != nil {
		return false, custom_errors.New(http.StatusInternalServerError, "服务器内部错误", err)
	}
	return !isNew, nil
}

// EnsureUserExists 确保用户存在
func (u *userService) EnsureUserExists(userAddress string) custom_errors.IAppError {
	if isNew, err := UserService.IsNewUser(userAddress); err != nil {
		return custom_errors.New(http.StatusInternalServerError, "服务器内部错误", err)
	} else if isNew {
		if err := dao.UserDao.CreateUser(userAddress); err != nil {
			return custom_errors.New(http.StatusInternalServerError, "服务器内部错误", err)
		}
	}
	return nil
}

func (u *userService) GetNonce(address string) (string, custom_errors.IAppError) {
	strCmd := configs.RedisClient.GetEx(appContext.NewTimeoutContext(), address, 3*time.Minute)
	//若当前查询得到的nonce还未过期，则直接返回，并更新过期时间
	if strCmd.Err() == nil {
		defer log.Printf(" 用户地址: %v; nonce: %v\n", address, strCmd.Val())
		return strCmd.Val(), nil
	} else if custom_errors.Is(strCmd.Err(), redis.Nil) {
		//已过期或不存在，执行生成nonce
		nonce := auth.GenerateNonce()

		if err := configs.RedisClient.SetEX(appContext.NewTimeoutContext(), address, nonce, 3*time.Minute).Err(); err == nil {
			defer log.Printf(" 用户地址是: %v; nonce: %v\n", address, nonce)
			return nonce, nil
		}
	}
	// Redis 查询出错或不可用
	defer log.Printf("查询 nonce 失败，用户地址: %v, 错误: %v", address, strCmd.Err())
	return "", custom_errors.New(http.StatusServiceUnavailable, "获取nonce失败", strCmd.Err())
}

func (u *userService) UpdateUser(toUpdate dto.UpdateUser) custom_errors.IAppError {
	if err := dao.UserDao.UpdateUser(toUpdate); err != nil {
		return custom_errors.New(http.StatusServiceUnavailable, "服务器内部错误", err)
	}
	return nil
}

func (u *userService) GetUserInfo(address string) (dto.UpdateUser, custom_errors.IAppError) {
	userInfo, err := dao.UserDao.GetUserInfo(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UpdateUser{}, custom_errors.New(http.StatusNotFound, "用户不存在", err)
		} else {
			return dto.UpdateUser{}, custom_errors.New(http.StatusInternalServerError, "服务器内部错误", err)
		}
	}
	return userInfo, nil
}
