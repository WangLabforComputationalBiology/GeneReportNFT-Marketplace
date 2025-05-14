package dao

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
	"sync"
)

var (
	gnftDao  *GNFTDao
	gnftOnce sync.Once
)

// GetGNFTDao 导出GNFTDao
func GetGNFTDao() *GNFTDao {
	gnftOnce.Do(func() {
		registerGNFTDao()
	})
	return gnftDao
}

type GNFTDao struct {
	db  *gorm.DB
	ctx context.Context
}

func registerGNFTDao() {

	gnftDao = &GNFTDao{
		db:  configs.DB,
		ctx: context.Background(),
	}
}

func (g *GNFTDao) DB() *gorm.DB {
	return g.db.WithContext(appContext.NewTimeoutContextByParent(g.ctx))
}
