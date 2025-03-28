package dao

import (
	"context"
	"gorm.io/gorm"
)

var orderDao *OrderDao

type OrderDao struct {
	db  *gorm.DB
	ctx context.Context
}

func RegisterOrderDao() {

	orderDao = &OrderDao{
		ctx: context.Background(),
		db:  DB,
	}
}

func (u *OrderDao) DB() *gorm.DB {
	return u.db.WithContext(u.ctx)
}
