package controllers

import (
	"github.com/gin-gonic/gin"
)

type GNFT struct{}

var (
	GNFTController = &GNFT{}
)

// GetGNFTInfoByOwner
//
//	@Summary		通过owner获取其拥有的GNFT信息
//	@Description	返回的信息包括集合数据以及集合中的所有GNFT
//	@Tags			collection管理
//	@Accept			json
//	@Produce		json
//	@Router			/collection/info/id [post]
//	@Param			Request	         body		dto.GetCollectionWithGNFTsByIDReq	true	"获取数据请求"
//	@Success		200				{object}	dto.Response	"成功获取集合信息"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		503				{object}	dto.ErrResponse	"获取集合失败"
func (g *GNFT) GetGNFTInfoByOwner(ctx *gin.Context) {

}
