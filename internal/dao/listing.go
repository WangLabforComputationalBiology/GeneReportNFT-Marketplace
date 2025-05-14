package dao

import (
	"GeneReport_platform/configs"
	"context"
	"gorm.io/gorm"
	"sync"
)

var (
	listingDao  *ListingDao
	listingOnce sync.Once
)

type ListingDao struct {
	db  *gorm.DB
	ctx context.Context
}

func (l *ListingDao) GetListingDao() *ListingDao {
	listingOnce.Do(func() {
		registerListingDao()
	})
	return listingDao
}
func registerListingDao() {

	listingDao = &ListingDao{
		ctx: context.Background(),
		db:  configs.DB,
	}
}

func (l *ListingDao) DB() *gorm.DB {
	return l.db.WithContext(l.ctx)
}
