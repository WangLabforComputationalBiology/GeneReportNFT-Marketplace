package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/pkg/CAPTCHA"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCATCHA 加载旋转式人机验证
func (s *Studio) GetCATCHA(ctx *gin.Context) {
	user, _ := ctx.Get("user_address")
	masterImage, thumbImage, err := CAPTCHA.GetRotateCAPTCHA(user.(string))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, dto.Response{Code: 200, Message: "加载人机验证成功", Data: gin.H{"master_image": masterImage, "thumb_image": thumbImage}})
}

// CheckCaptcha 执行旋转式人机验证
func (s *Studio) CheckCaptcha(ctx *gin.Context) {
	user, _ := ctx.Get("user_address")
	var req dto.CheckCAPTCHAReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重试",
		})
		return
	}
	isPass, err := CAPTCHA.CheckRotateCAPTCHA(user.(string), req.Angle)
	if err != nil {
		ctx.Error(err)
		return
	}
	if isPass {
		ctx.JSON(200, dto.Response{Code: 200, Message: "人机验证成功"})
		return
	} else {
		ctx.JSON(200, dto.Response{Code: 200, Message: "人机验证失败"})
		return
	}
}
