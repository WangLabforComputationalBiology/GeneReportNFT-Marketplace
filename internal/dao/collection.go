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
	collectionDao  *CollectionDao
	collectionOnce sync.Once
)

type CollectionDao struct {
	db  *gorm.DB
	ctx context.Context
}
type CollectionWithGNFT struct {
	Collection models.Collection
	GNFT       models.GNFT
}

func GetCollectionDao() *CollectionDao {
	collectionOnce.Do(func() {
		collectionDao = &CollectionDao{
			db:  configs.DB,
			ctx: context.Background(),
		}
	})
	return collectionDao
}

func (c *CollectionDao) DB() *gorm.DB {
	return c.db.WithContext(appContext.NewTimeoutContextByParent(c.ctx))
}

// GetCollectionWithGNFTByID 获取collectionID的collections与gnfts联表结果，并以CollectionWithGNFT作为联表单元行返回
func (c *CollectionDao) GetCollectionWithGNFTByID(collectionID string) (results []CollectionWithGNFT, err error) {
	err = c.DB().Select("collections.*, gnfts.*").
		Table("collections").
		Joins("LEFT JOIN gnfts ON gnfts.collection_id = collections.id").
		Where("collections.id = ?", collectionID).
		Order("gnfts.created_at desc").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetCollectionWithGNFTByCreator 获取creator的collections与gnfts联表结果，并以CollectionWithGNFT作为联表单元行返回
func (c *CollectionDao) GetCollectionWithGNFTByCreator(creator string) (results []CollectionWithGNFT, err error) {
	err = c.DB().Select("collections.*, gnfts.*").
		Table("collections").
		Joins("LEFT JOIN gnfts ON gnfts.collection_id = collections.id").
		Where("collections.creator = ?", creator).
		Group("collections.id").
		Order("gnfts.created_at desc").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetCollectionsInfoByCreator 获取Creator创建的所有collection信息
func (c *CollectionDao) GetCollectionsInfoByCreator(creator string) (results []models.Collection, err error) {
	err = c.DB().Select("collections.*").
		Table("collections").
		Where("creator = ?", creator).
		Order("created_at desc").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
