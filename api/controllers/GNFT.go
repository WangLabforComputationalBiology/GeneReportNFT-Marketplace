package controllers

import "github.com/gin-gonic/gin"

type GNFTController struct{}

var (
	GNFTCtrller = &GNFTController{}
)

func (cGNFT *GNFTController) GetInfo(ctx *gin.Context) {

}

func (cGNFT *GNFTController) GetList(ctx *gin.Context) {

}

func (cGNFT *GNFTController) Mint(ctx *gin.Context) {

}

// Burn 销毁GNFT
func (cGNFT *GNFTController) Burn(ctx *gin.Context) {

}

// PutOnSale 上架GNFT
func (cGNFT *GNFTController) PutOnSale() string {
	// 添加逻辑代码
	return "PutOnSale"
}

// TakeOffSale 下架GNFT
func (cGNFT *GNFTController) TakeOffSale() string {
	// 添加逻辑代码
	return "TakeOffSale"
}
