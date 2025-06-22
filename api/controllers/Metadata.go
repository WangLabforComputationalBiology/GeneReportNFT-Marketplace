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

// GetMetadataInfoByOwner 获取指定owner的metadata信息
func (m *Metadata) GetMetadataInfoByOwner(ctx *gin.Context) {
	var req dto.GetMetadataByOwnerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	gnftInfos, err := services.MetadataServ.GetMetadataOverviewByOwner(req.Owner)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "成功获取GNFT信息",
		Data:    gnftInfos,
	})
	return
}

// GetMetadataInfoByDataHash 获取指定id的metadata信息
func (m *Metadata) GetMetadataInfoByDataHash(ctx *gin.Context) {
	var req dto.GetMetadataByOwnerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}

}

func (m *Metadata) GetGeneTypeDataByID(ctx *gin.Context) {
	services.GetDataImpl(ctx)
}

// GetData 获取geneType数据
func (m *Metadata) GetData(ctx *gin.Context) {
	services.GetDataImpl(ctx)
}
