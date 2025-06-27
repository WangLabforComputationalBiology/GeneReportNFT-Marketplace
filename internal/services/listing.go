package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/pkg/appErrors"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

var (
	ListingServ *ListingService
)

type ListingService struct {
	iListingBase
}

// 订单基础接口
type iListingBase interface {
}

func RegisterOrderService() {
	ListingServ = &ListingService{}
}

/*fill your method here*/

// GetDataImpl 获取GeneType数据
func GetDataImpl(profileId, category string) (map[string]interface{}, error) {
	unique := dto.UniqueProfiles{}
	//在数据库查出profileId一样记录
	if profileId != "" {
		configs.DB.Where("profile_id = ?", profileId).Find(&unique)
	}
	if unique.Status == 0 {
		return nil, appErrors.New(http.StatusOK, "数据还没处理完成，请稍等!")
	}

	t, ok := dto.GetStructType(category)
	if ok {
		fmt.Println("Struct Type:", t)
		instance := reflect.New(t).Interface()
		fmt.Println("Instance:", instance)
	}

	// 创建该类型的切片用于存储查询结果
	sliceType := reflect.SliceOf(t)

	err := configs.DB.Model(t).Where("profile_id = ?", profileId).Find(sliceType).Error
	if err != nil {
		return nil, appErrors.New(http.StatusInternalServerError, "查询失败", err)
	}

	var results map[string]interface{}
	tempResults, _ := json.Marshal(sliceType)
	_ = json.Unmarshal(tempResults, &results)
	return results, nil
}
