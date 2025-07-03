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
	activityDao  *ActivityDao
	activityOnce sync.Once
)

type ActivityDao struct {
	db  *gorm.DB
	ctx context.Context
}

// GetActivityDao 导出ActivityDao
func GetActivityDao() *ActivityDao {
	activityOnce.Do(func() {
		activityDao = &ActivityDao{
			db:  configs.DB,
			ctx: context.Background(),
		}
	})
	return activityDao
}

func (a *ActivityDao) DB() *gorm.DB {
	return a.db.WithContext(appContext.NewTimeoutContextByParent(a.ctx))
}

func (a *ActivityDao) GetLatestViewAccess(userAddress string, dataHash string) (results *models.Activity, err error) {
	err = a.DB().Select("activities.*").
		Table("activities").
		Where("event='addViewAccess' AND user_address = ? AND metadata = ?", userAddress, dataHash).
		Order("time desc").
		First(&results).Error
	return results, err
}

func (a *ActivityDao) NewViewAccess(activity *models.Activity) (err error) {
	err = a.DB().Table("activities").
		Create(&activity).Error
	return err
}
