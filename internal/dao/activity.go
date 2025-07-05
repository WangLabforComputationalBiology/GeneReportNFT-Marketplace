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
		Where("event IN ? AND user_address = ? AND metadata = ?", []string{"NewViewAccess", "RenewalViewAccess"}, userAddress, dataHash).
		Order("time desc").
		First(&results).Error
	return results, err
}

func (a *ActivityDao) NewViewAccess(activity *models.Activity) (err error) {
	activity.Event = "NewViewAccess"
	err = a.DB().Table("activities").
		Create(&activity).Error
	return err
}

func (a *ActivityDao) RenewalViewAccess(activity *models.Activity) (err error) {
	activity.Event = "RenewalViewAccess"
	err = a.DB().Table("activities").
		Create(&activity).Error
	return err
}

func (a *ActivityDao) IsViewAccessExist(dataHash, viewer string) (bool, error) {
	var count int
	err := a.DB().Select("count(*)").
		Table("activities").
		Where("event IN ? AND user_address = ? AND metadata = ?", []string{"NewViewAccess", "RenewalViewAccess"}, viewer, dataHash).
		Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
