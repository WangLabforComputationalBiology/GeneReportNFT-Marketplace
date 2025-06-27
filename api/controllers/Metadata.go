package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Metadata struct{}

var (
	MetadataController = &Metadata{}
)

// GetMetadataOverviewByOwner 获取指定owner的metadata信息
func (m *Metadata) GetMetadataOverviewByOwner(ctx *gin.Context) {
	var req dto.GetMetadataOverviewByOwnerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	results, err := services.MetadataServ.GetMetadataOverviewByOwner(req.Owner)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    results,
	})
	return
}

// GetMetadataDetailByDataHash 获取指定数据哈希的metadata详情
func (m *Metadata) GetMetadataDetailByDataHash(ctx *gin.Context) {
	dataHash := ctx.Param("dataHash")
	if dataHash == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误",
		})
		return
	}
	toResp, err := services.MetadataServ.GetMetadataDetailByDataHash(dataHash)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toResp,
	})
}

func (m *Metadata) GetAllMetadataOverview(ctx *gin.Context) {
	multiMetadata, err := services.MetadataServ.GetAllMetadataOverview()
	if err != nil {
		ctx.Error(err)
		return
	}
	//成功
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    multiMetadata,
	})
}
