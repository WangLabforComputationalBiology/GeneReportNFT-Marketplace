package dao

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
	"time"
)

var UserDao *userDao

type userDao struct {
	db  *gorm.DB
	ctx context.Context
}

func RegisterUserDao() {
	UserDao = &userDao{
		ctx: context.Background(),
		db:  configs.DB,
	}
}

func (u *userDao) DB() *gorm.DB {
	return u.db.WithContext(appContext.NewTimeoutContextByParent(u.ctx))
}

func (u *userDao) IsExist(address string) (bool, error) {
	var count int64
	if err := u.DB().Model(&models.User{}).Where("address = ?", address).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *userDao) CreateUser(address string) error {
	return u.DB().Model(&models.User{}).Create(&models.User{Address: address, Name: "unnamed", CreateAt: time.Now()}).Error
}

func (u *userDao) UpdateUser(toUpdate dto.UpdateUser) error {
	return u.DB().Model(&models.User{}).Model(&models.User{}).Where("address = ?", toUpdate.Address).Updates(&models.User{Name: toUpdate.Name}).Error
}

func (u *userDao) GetUser(address string) (models.User, error) {
	var userInfo models.User
	return userInfo, u.DB().Model(&models.User{}).Where("address = ?", address).First(&userInfo).Error
}
