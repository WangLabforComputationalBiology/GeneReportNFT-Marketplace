package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/dao"
	"GeneReport_platform/pkg/appErrors"
	"fmt"
	"github.com/gin-gonic/gin"
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

func GetDataImpl(ctx *gin.Context) {
	//获取参数
	param := ctx.Param("param")
	t, ok := dto.GetStructType(param)
	if ok {
		fmt.Println("Struct Type:", t)
		instance := reflect.New(t).Interface()
		fmt.Println("Instance:", instance)
	}

	// 创建该类型的切片用于存储查询结果
	sliceType := reflect.SliceOf(t)
	result := reflect.New(sliceType).Interface()

	// 使用 GORM 查询数据并填充到 result 中
	pageSize := 5
	pageNum := 1
	err := configs.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(result).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, result)

}

func (s ListingService) GetListings(collection_id, identifier string) (dto.GetListingsResp, error) {
	targetListings, err := dao.GetListingDao().GetListings(collection_id, identifier)
	if err != nil {
		return dto.GetListingsResp{}, appErrors.New(503, "获取上架单失败", err)
	}
	return dto.GetListingsResp{Listings: targetListings}, nil
}
