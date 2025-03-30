package dao

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
)

var gnftDao *GNFTDao

type GNFTDao struct {
	db  *gorm.DB
	ctx context.Context
}

func RegisterGNFTDao() {

	gnftDao = &GNFTDao{
		db:  configs.DB,
		ctx: context.Background(),
	}
}

func (u *GNFTDao) DB() *gorm.DB {
	return u.db.WithContext(appContext.NewTimeoutContextByParent(u.ctx))
}
