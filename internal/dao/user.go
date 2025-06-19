package dao

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	userDao  *UserDao
	userOnce sync.Once
)

type UserDao struct {
	db  *gorm.DB
	ctx context.Context
}

// UpdateUser userModel for update
type UpdateUser struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
}

func GetUserDao() *UserDao {
	userOnce.Do(func() {
		registerUserDao()
	})
	return userDao
}

func registerUserDao() {
	userDao = &UserDao{
		ctx: context.Background(),
		db:  configs.DB,
	}
}

func (u *UserDao) DB() *gorm.DB {
	return u.db.WithContext(appContext.NewTimeoutContextByParent(u.ctx))
}

func (u *UserDao) IsExist(address string) (bool, error) {
	var count int64
	if err := u.DB().Model(&models.User{}).Where("address = ?", address).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *UserDao) CreateUser(address string) error {
	return u.DB().Model(&models.User{}).Create(&models.User{Address: address, Name: "unnamed", CreateAt: time.Now()}).Error
}

func (u *UserDao) UpdateUser(toUpdate UpdateUser) error {
	return u.DB().Model(&models.User{}).Where("address = ?", toUpdate.Address).Updates(&models.User{Name: toUpdate.Name}).Error
}

func (u *UserDao) GetUser(address string) (models.User, error) {
	var userInfo models.User
	return userInfo, u.DB().Model(&models.User{}).Where("address = ?", address).First(&userInfo).Error
}

func (u *UserDao) GetEmailSuffix(institutionName string) (results string, err error) {
	err = u.DB().Select("institutions.email").
		Where("institutions.name = ?", institutionName).
		Scan(&results).Error
	return results, err
}
