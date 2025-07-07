package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GeneSharing struct{}

var (
	GeneSharingController = &GeneSharing{}
)

// GetGeneSharingDetailsByContractAddress 通过合约地址获取GeneSharing合集详情
func (g *GeneSharing) GetGeneSharingDetailsByContractAddress(ctx *gin.Context) {
	var req dto.GetGeneSharingDetailsByContractAddressReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	toResp, err := services.GeneSharingServ.GetGeneSharingDetailsByContractAddress(req.GeneSharingContractAddress)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toResp,
	})
	return
}

// GetGeneSharingByCreator 获取用户创建的基因共享集合
func (g *GeneSharing) GetGeneSharingByCreator(ctx *gin.Context) {
	var req dto.GetGeneSharingOverviewByCreatorReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	toResp, err := services.GeneSharingServ.GetGeneSharingOverviewByCreator(req.CreatorAddress)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    toResp,
	})
	return
}

// GetAllGeneSharingOverview 获取plaza主页（GeneSharing概览）
func (g *GeneSharing) GetAllGeneSharingOverview(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数格式错误",
		})
		return
	}

	toResp, err := services.GeneSharingServ.GetAllGeneSharingOverview(page)
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
	return
}
