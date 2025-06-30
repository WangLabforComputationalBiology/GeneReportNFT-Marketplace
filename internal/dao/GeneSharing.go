package dao

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"GeneReport_platform/pkg/appContext"
	"context"
	"gorm.io/gorm"
	"sync"
)

var (
	geneSharingDao  *GeneSharingDao
	geneSharingOnce sync.Once
)

type GeneSharingDetail struct {
	GeneSharing       models.GeneSharing
	MetadataOverviews []dto.MetadataOverview `json:"metadata_overviews"` //该GeneSharing集合下的所有Metadata的overview
}

type GeneSharingDao struct {
	db  *gorm.DB
	ctx context.Context
}

// GetGeneSharingDao 导出GeneSharingDao
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

// GetGeneSharingDetailsByAddress 根据地址获取GeneSharing详情
func (g *GeneSharingDao) GetGeneSharingDetailsByAddress(geneSharingContractAddress string) (results GeneSharingDetail, err error) {
	type tempResult struct {
		GeneSharing       models.GeneSharing
		MetadataOverviews dto.MetadataOverview
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

// GetGeneSharingOverviewByCreator 获取creator的geneSharing的概览信息
func (g *GeneSharingDao) GetGeneSharingOverviewByCreator(creator string) (results []dto.GeneSharingOverview, err error) {
	err = g.DB().Select("geneSharings.*").
		Where("geneSharings.creator = ?", creator).
		Order("geneSharings.created_at desc").
		Scan(&results).Error

	return results, nil
}

// CreateGeneSharing 创建geneSharing
func (g *GeneSharingDao) CreateGeneSharing(geneSharing *models.GeneSharing) error {
	err := g.DB().Table("geneSharings").Create(&geneSharing).Error
	return err
}
