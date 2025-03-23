package dao

import (
	"GeneReport_platform/internal/dao/global"
	"context"
	"gorm.io/gorm"
)

var gnftDao *GNFTDao

type GNFTDao struct {
	table string
	db    *gorm.DB
	ctx   context.Context
}

func RegisterGNFTDao() {

	gnftDao = &GNFTDao{
		ctx:   global.Ctx,
		db:    global.DB,
		table: "gnft",
	}
}

func (u *GNFTDao) DB() *gorm.DB {
	return u.db.WithContext(u.ctx)
}
