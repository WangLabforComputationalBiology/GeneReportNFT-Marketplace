package dao

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
	"sync"
)

var (
	gnftDao  *GNFTDao
	gnftOnce sync.Once
)

type GNFTWithCollection struct {
	GNFT       models.GNFT
	Collection models.Collection
	Quantity   int `gorm:"column:quantity" json:"quantity"`
}

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

func (g *GNFTDao) GetGNFTInfosByOwner(owner string) (results []GNFTWithCollection, err error) {
	err = g.DB().Select("collections.*, gnfts.*,ownerships.quantity").
		Table("(SELECT collection_id, identifier, quantity from  ownerships WHERE owner = ?) AS ownership", owner).
		Joins("JOIN gnfts ON  gnfts.collection_id = ownership.collection_id AND gnfts.identifier = ownership.identifier").
		Joins("JOIN collections ON collections.id = gnfts.collection_id").
		Order("gnfts.collection_id, gnfts.identifier").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
