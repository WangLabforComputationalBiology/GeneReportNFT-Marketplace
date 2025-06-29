package dao

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/models"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"sync"
)

var (
	metadataDao  *Metadata
	metadataOnce sync.Once
)

// GetMetadataDao 导出MetadataDao
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
	return m.db
}

// GetMetadataOverviewByDataHash 通过data_hash获取metadata
func (m *Metadata) GetMetadataOverviewByDataHash(dataHashBase64 string) (result models.Metadata, err error) {
	err = m.DB().Select("metadatas.*").
		Table("metadatas").
		Where("metadatas.data_hash = ?&& metadatas.is_hidden = 0", dataHashBase64).
		Order("metadatas.category desc").
		Scan(&result).Error
	if err != nil {
		return models.Metadata{}, err
	}
	return result, nil
}

// GetMetadataOverviewByOwner 通过Owner获取metadata信息
func (m *Metadata) GetMetadataOverviewByOwner(owner string) (results []dto.MetadataOverview, err error) {
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

func (m *Metadata) GetAllMetadataOverview() (results []dto.MetadataOverview, err error) {
	err = m.DB().Select("metadatas.*").
		Table("metadatas").
		Where("metadatas.is_hidden = 0").
		Order("metadatas.created_at desc").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// CompleteMetadataInfo 补全预构建的数据库的Metadata信息
func (m *Metadata) CompleteMetadataInfo(profileID string, toUpdate dto.UpdateMetadata) error {
	// 转为 map[string]interface{}
	updates := make(map[string]interface{})
	data, err := json.Marshal(toUpdate)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &updates); err != nil {
		return err
	}

	// 开启事务
	tx := m.DB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新所有匹配 profileID 的记录
	err = tx.Table("metadatas").
		Where("profile_id = ? AND is_hidden = 0", profileID).
		Updates(updates).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m *Metadata) GetGenoType(profileId, category string) (results []models.GenoType, err error) {
	err = m.DB().Select("genotypes.*").
		Table("genotypes").
		Where("profile_id = ? AND type = ?", profileId, category).
		Order("report_id asc").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, err
}
