package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GNFT struct{}

var (
	GNFTController = &GNFT{}
)

// GetGNFTInfoByOwner
//
//	@Summary		通过owner获取其拥有的GNFT信息
//	@Description	返回的信息包括集合数据以及集合中的所有GNFT
//	@Tags			GNFT管理
//	@Accept			json
//	@Produce		json
//	@Router			/gnft/info/owner [post]
//	@Param			Request	         body		dto.GetGNFTInfosByOwnerReq	true	"获取数据请求"
//	@Success		200				{object}	dto.Response	"成功获取集合信息"
//	@Failure		400				{object}	dto.ErrResponse	"请求体格式错误"
//	@Failure		503				{object}	dto.ErrResponse	"获取集合失败"
func (g *GNFT) GetGNFTInfoByOwner(ctx *gin.Context) {
	var req dto.GetGNFTInfosByOwnerReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	gnftInfos, err := services.GNFTServ.GetGNFTInfosByOwner(req.Owner)
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
