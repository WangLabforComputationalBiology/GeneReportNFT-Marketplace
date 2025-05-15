package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Market struct{}

var (
	MarketController = &Market{}
)

func (l *Market) GetListings(ctx *gin.Context) {
	var req dto.GetListingsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重新登录",
		})
		return
	}
	listings, err := services.ListingServ.GetListings(req.CollectionId, req.Identifier)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "获取GNFT上架单成功",
		Data:    listings,
	})
}
