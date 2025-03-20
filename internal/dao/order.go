package dao

import (
	"GeneReport_platform/internal/dao/global"
	"context"
	"gorm.io/gorm"
)

var orderDao *OrderDao

type OrderDao struct {
	table string
	db    *gorm.DB
	ctx   context.Context
}

func RegisterOrderDao() {

	orderDao = &OrderDao{
		ctx:   global.Ctx,
		db:    global.DB,
		table: "order_dto",
	}
}

func (u *OrderDao) DB() *gorm.DB {
	return u.db.WithContext(u.ctx)
}
