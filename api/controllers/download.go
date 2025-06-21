package controllers

import (
	"github.com/gin-gonic/gin"
)

type Download struct{}

var (
	DownloadController = &Download{}
)

func (d *Download) DownloadFile(ctx *gin.Context) {
	ctx.Param("shortCode")

}
