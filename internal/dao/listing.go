package dao

import (
	"GeneReport_platform/configs"
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
		db:  configs.DB,
	}
}

func (u *OrderDao) DB() *gorm.DB {
	return u.db.WithContext(u.ctx)
}
