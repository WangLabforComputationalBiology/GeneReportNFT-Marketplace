package dao

import (
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
)

var gnftDao *GNFTDao

type GNFTDao struct {
	db        *gorm.DB
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func RegisterGNFTDao() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	gnftDao = &GNFTDao{
		db:        DB,
		ctx:       ctx,
		ctxCancel: cancelFunc,
	}
}

func (u *GNFTDao) DB() *gorm.DB {
	return u.db.WithContext(appContext.NewTimeoutContextByParent(u.ctx))
}
