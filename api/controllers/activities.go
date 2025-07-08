package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
)

type Activity struct{}

var (
	ActivityController = &Activity{}
)

func (a *Activity) GetTraceDataByDataHash(ctx *gin.Context) {
	toResp, err := services.MetadataServ.GetTraceData(ctx.Query("data_hash"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	//成功
	ctx.JSON(200, dto.Response{
		Code:    200,
		Message: "success",
		Data:    toResp,
	})
}

func (a *Activity) GetActivityByUser(ctx *gin.Context) {
	toResp, err := services.UserServ.GetActivityByUser(ctx.GetString("user_address"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	//成功
	ctx.JSON(200, dto.Response{
		Code:    200,
		Message: "success",
		Data:    toResp,
	})
}
