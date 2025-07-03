package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Download struct{}

var (
	DownloadController = &Download{}
)

func (d *Download) DownloadFile(ctx *gin.Context) {
	shortCode := ctx.Param("short_code")
	if shortCode == "" {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数缺失",
		})
		return
	}
	pr, err := services.DownloadServ.DownloadFile(shortCode, ctx.GetString("user_address"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Header("Content-Type", "application/zip")
	ctx.Header("Content-Disposition", "attachment; filename=example.zip")
	ctx.Stream(func(w io.Writer) bool {
		_, err := io.Copy(w, pr)
		if err != nil {
			return false
		}
		return false // 流完成
	})

}
