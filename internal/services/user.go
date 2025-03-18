package services

import (
	"GeneReport_platform/api/dto/user_dto"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/pkg/auth"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
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
func (u *userService) IsNewUser(userAddress string) (bool, error) {
	isNew, err := dao.UserDao.IsExist(userAddress)
	if err != nil {
		return false, err
	}
	return !isNew, nil
}

// EnsureUserExists 确保用户存在
func (u *userService) EnsureUserExists(userAddress string) error {
	if isNew, err := UserService.IsNewUser(userAddress); err != nil {
		return err
	} else if isNew {
		return dao.UserDao.CreateUser(userAddress)
	} else {
		return nil
	}

}
func (u *userService) GetNonce(address string) (string, error) {
	strCmd := global.RedisClient.GetEx(global.Ctx, address, 3*time.Minute)
	//若当前查询得到的nonce还未过期，则直接返回，并更新过期时间
	if strCmd.Err() == nil {
		defer log.Printf(" 用户地址: %v; nonce: %v\n", address, strCmd.Val())
		return strCmd.Val(), nil
	} else if errors.Is(strCmd.Err(), redis.Nil) {
		//已过期或不存在，执行生成nonce
		nonce := auth.GenerateNonce()

		if err := global.RedisClient.SetEX(global.Ctx, address, nonce, 3*time.Minute).Err(); err == nil {
			defer log.Printf(" 用户地址是: %v; nonce: %v\n", address, nonce)
			return nonce, nil
		}
	}
	// Redis 查询出错或不可用
	defer log.Printf("查询 nonce 失败，用户地址: %v, 错误: %v", address, strCmd.Err())
	return "", strCmd.Err()
}

func (u *userService) UpdateUser(toUpdate user_dto.UpdateUser) error {
	return dao.UserDao.UpdateUser(toUpdate)
}

func (u *userService) GetUserInfo(address string) (user_dto.UpdateUser, error) {
	userInfo, err := dao.UserDao.GetUserInfo(address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user_dto.UpdateUser{}, errors.New("用户不存在")
		} else {
			return user_dto.UpdateUser{}, err
		}
	}
	return userInfo, nil
}
