package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
			Message: "请求体格式错误",
		})
		return
	}
	results, err := services.MetadataServ.GetMetadataOverviewByOwner(req.Owner)
	if err != nil {
		_ = ctx.Error(err)
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
	dataHash := ctx.Param("data_hash")
	if dataHash == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数缺失",
		})
		return
	}
	toResp, err := services.MetadataServ.GetMetadataDetailByDataHash(dataHash)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toResp,
	})
}

// GetAllMetadataOverview 获取plaza主页（Metadata概览）
func (m *Metadata) GetAllMetadataOverview(ctx *gin.Context) {

	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数格式错误",
		})
		return
	}

	toResp, err := services.MetadataServ.GetAllMetadataOverview(page)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	//成功
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toResp,
	})
}

func (m *Metadata) GetGenoTypeZip(ctx *gin.Context) {
	dataHash := ctx.Param("data_hash")
	if dataHash == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数缺失",
		})
		return
	}

	toResp, err := services.MetadataServ.GetGenoTypeZip(dataHash, ctx.GetString("user_address"), ctx.GetString("pub_key"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toResp,
	})

}

func (m *Metadata) ObtainViewAccess(ctx *gin.Context) {
	var req dto.NewViewAccessReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}

	toResp, err := services.MetadataServ.ObtainViewAccess(req.TxHash, ctx.GetString("user_address"), ctx.GetString("pub_key"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toResp,
	})

}
