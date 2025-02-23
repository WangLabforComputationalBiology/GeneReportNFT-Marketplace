package dao

import (
	"GeneReport_platform/internal/dao/global"
	"GeneReport_platform/internal/models/entity"
	"context"
	"gorm.io/gorm"
)

var userDao *UserDao

type UserDao struct {
	table string
	db    *gorm.DB
	ctx   context.Context
}

func RegisterUserDao() {
	userDao = &UserDao{
		ctx:   global.Ctx,
		db:    global.DB,
		table: "user",
	}
}

func (u *UserDao) DB() *gorm.DB {
	return u.db.WithContext(u.ctx).Model(&entity.User{})
}
