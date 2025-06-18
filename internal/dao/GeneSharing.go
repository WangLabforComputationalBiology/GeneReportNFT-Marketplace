package dao

import (
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	geneSharingDao  *GeneSharingDao
	geneSharingOnce sync.Once
)

type GeneSharingOverview struct {
	ID          string    `gorm:"primaryKey;type:varchar(32)" json:"ID"`             //以profile_id作为id
	Name        string    `gorm:"type:varchar(32)" json:"name"`                      //sharing集合名称
	Description string    `gorm:"type:text" json:"description"`                      //sharing集合描述
	Creator     string    `gorm:"type:varchar(42);index:idx_creator" json:"creator"` //sharing集合创建者
	CreatedAt   time.Time `gorm:"type:datetime" json:"created_at"`                   //sharing集合创建时间
	ItemAmount  int       `gorm:"type:int;" json:"item_amount"`                      //MetaData数量
	IsOfficial  bool      `gorm:"type:tinyint(1);" json:"is_official"`               //是否第三方官方授权构建
	Tags        string    `gorm:"type:varchar(255)" json:"tags"`                     //标签，以分号分隔，,ex:third party:wegene;...
}

type GeneSharingOverviews struct {
	GeneSharings []GeneSharingOverview `json:"geneSharings"`
}

type GeneSharingDetail struct {
	GeneSharing       models.GeneSharing
	MetadataOverviews []MetadataOverview `json:"metadata_overviews"` //该GeneSharing集合下的所有Metadata的overview
}

type GeneSharingDao struct {
	db  *gorm.DB
	ctx context.Context
}

func GetGeneSharingDao() *GeneSharingDao {
	geneSharingOnce.Do(func() {
		geneSharingDao = &GeneSharingDao{
			db:  configs.DB,
			ctx: context.Background(),
		}
	})
	return geneSharingDao
}

func (g *GeneSharingDao) DB() *gorm.DB {
	return g.db.WithContext(appContext.NewTimeoutContextByParent(g.ctx))
}

// GetGeneSharingDetailsByAddress 获取，并以CollectionWithGNFT作为联表单元行返回
func (g *GeneSharingDao) GetGeneSharingDetailsByAddress(geneSharingContractAddress string) (results GeneSharingDetail, err error) {
	type tempResult struct {
		GeneSharing       models.GeneSharing
		MetadataOverviews MetadataOverview
	}
	var tempResults []tempResult

	// 查询
	err = g.DB().Select("geneSharings.*,"+"metadatas.data_hash+','+metadatas.format+','+metadatas.sex+','+metadatas.category+','+metadatas.name+','+metadatas.is_sharable+','+metadatas.created_at").
		Table("geneSharings").
		Joins("INNER JOIN geneSharing_metadatas ON geneSharings.contract_address = geneSharing_metadatas.geneSharing_contract_address").
		Joins("INNER JOIN metadatas ON geneSharing_metadatas.metadata_hash = metadatas.data_hash").
		Where("geneSharings.contract_address = ? AND metadatas.is_hidden = 0", geneSharingContractAddress).
		Order("geneSharings.created_at desc").
		Scan(&tempResults).Error
	if err != nil {
		return GeneSharingDetail{}, err
	}

	results = GeneSharingDetail{
		GeneSharing: tempResults[0].GeneSharing,
	}
	//循环添加元素
	for _, tempResult := range tempResults {
		results.MetadataOverviews = append(results.MetadataOverviews, tempResult.MetadataOverviews)
	}
	return results, nil
}

// GetGeneSharingOverviewByCreator 获取creator的collections与gnfts联表结果，并以CollectionWithGNFT作为联表单元行返回
func (g *GeneSharingDao) GetGeneSharingOverviewByCreator(creator string) (results []GeneSharingOverview, err error) {
	err = g.DB().Select("geneSharings.*").
		Where("geneSharings.creator = ?", creator).
		Order("geneSharings.created_at desc").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}
