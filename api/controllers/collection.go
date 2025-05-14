package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Collection struct{}

var (
	CollectionController = &Collection{}
)

// GetCollectionInfoByID
//
//	@Summary		通过collection_id获取集合信息
//	@Description	返回的信息包括集合数据以及集合中的所有GNFT
//	@Tags			collection管理
//	@Accept			json
//	@Produce		json
//	@Router			/collection/info [post]
//	@Body			{collection_id} dto.GetCollectionInfoReq
//	@Success		200				{object}	dto.Response	"成功获取集合信息"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		503				{object}	dto.ErrResponse	"获取集合失败"
func GetCollectionInfoByID(ctx *gin.Context) {
	var req dto.GetCollectionWithGNFTsByIDReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	collectionInfoRes, err := services.CollectionServ.GetCollectionInfoByID(req.CollectionID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "成功获取集合信息",
		Data:    collectionInfoRes,
	})
	return
}

// GetCollectionInfoByCreator
//
//	@Summary		通过creator所创建的所有集合信息
//	@Description	返回的信息包括集合数据以及集合中的所有GNFT
//	@Tags			collection管理
//	@Accept			json
//	@Produce		json
//	@Router			/collection/info [post]
//	@Body			{collection_id} dto.GetCollectionsWithGNFTsByCreatorReq
//	@Success		200				{object}	dto.Response	"成功获取集合信息"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		503				{object}	dto.ErrResponse	"获取集合失败"
func GetCollectionInfoByCreator(ctx *gin.Context) {
	var req dto.GetCollectionsWithGNFTsByCreatorReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	collectionInfosResp, err := services.CollectionServ.GetCollectionInfosByCreator(req.Creator)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "成功获取集合信息",
		Data:    collectionInfosResp,
	})
	return
}
