package dao

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
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

func GetListingDao() *ListingDao {
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

func (l *ListingDao) GetListings(collectionId string, identifier string) (listings []models.Listing, err error) {
	err = l.DB().Select("*").
		Table("listings").
		Where("collection_id = ? and identifier = ?", collectionId, identifier).
		Order("sale_amount asc").
		Find(&listings).Error
	if err != nil {
		return nil, err
	}
	return listings, nil
}
