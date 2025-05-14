package controllers

import (
	"github.com/gin-gonic/gin"
)

type GNFT struct{}

var (
	GNFTController = &GNFT{}
)

func (cGNFT *GNFT) GetGNFTInfo(ctx *gin.Context) {

}

func (cGNFT *GNFT) GetListings(ctx *gin.Context) {

}
