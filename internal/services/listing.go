package services

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
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

// GetDataImpl 获取GeneType数据
func GetDataImpl(ctx *gin.Context) {
	address := ctx.GetString("user_address")
	profileId := ctx.Query("profileId")
	unique := dto.UniqueProfiles{}
	//在数据库查出profileId一样记录
	if profileId != "" {
		configs.DB.Where("profile_id = ?", profileId).Find(&unique)
	}
	if unique.Status == 0 {
		ctx.JSON(http.StatusOK, gin.H{"msg": "数据还没处理完成，请稍等!"})
		return
	}

	//先将访问数据存到数据库中
	record := dto.DataVisitRecord{
		Address:   address,
		ProfileID: profileId,
	}
	configs.DB.Create(&record)
	//todo 再将数据拼接成string存到链上
	//

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
