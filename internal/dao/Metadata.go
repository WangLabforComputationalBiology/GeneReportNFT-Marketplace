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
	metadataDao  *Metadata
	metadataOnce sync.Once
)

type MetadataOverview struct {
	DataHash   string    `gorm:"primaryKey;type:varchar(36)" json:"data_hash"`
	Format     string    `gorm:"type:varchar(32)" json:"format"`
	Sex        bool      `gorm:"type:varchar(32)" json:"sex"`
	Category   string    `gorm:"type:varchar(32)" json:"category"`
	Name       string    `gorm:"index:idx_name;type:varchar(32)" json:"name"`
	IsSharable bool      `gorm:"type:tinyint(1)" json:"is_sharable"`
	CreatedAt  time.Time `gorm:"type:datetime" json:"created_at"`
}

type MetadataWithGeneSharing struct {
	Metadata    models.Metadata
	GeneSharing models.GeneSharing
}

// GetMetadataDao 导出GNFTDao
func GetMetadataDao() *Metadata {
	metadataOnce.Do(func() {
		registerMetadataDao()
	})
	return metadataDao
}

type Metadata struct {
	db  *gorm.DB
	ctx context.Context
}

func registerMetadataDao() {
	metadataDao = &Metadata{
		db:  configs.DB,
		ctx: context.Background(),
	}
}

func (m *Metadata) DB() *gorm.DB {
	return m.db.WithContext(appContext.NewTimeoutContextByParent(m.ctx))
}

// GetMetadataDetailByDataHash 通过data_hash获取metadata
func (m *Metadata) GetMetadataDetailByDataHash(dataHashBase64 string) (results models.Metadata, err error) {
	err = m.DB().Select("metadatas.*").
		Where("metadatas.data_hash = ?&& metadatas.is_hidden = 0", dataHashBase64).
		Order("metadatas.category desc").
		Scan(results).Error
	if err != nil {
		return models.Metadata{}, err
	}
	return results, nil
}

// GetMetadataOverviewByOwner 通过Owner获取metadata信息
func (m *Metadata) GetMetadataOverviewByOwner(owner string) (results []MetadataOverview, err error) {
	err = m.DB().Select("metadatas.*").
		Where("metadatas.owner = ? && metadatas.is_hidden = 0", owner).
		Group("metadatas.contract_address").
		Order("metadatas.category desc").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// GetMetadataDetailByProfileId 通过profileID获取预构建的metadata信息
func (m *Metadata) GetMetadataDetailByProfileId(profileID string) (results []models.Metadata, err error) {
	err = m.DB().Select("metadatas.*").
		Table("metadatas").
		Where("metadatas.profile_id = ? && metadatas.is_hidden = 0", profileID).
		Order("metadatas.category desc").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
