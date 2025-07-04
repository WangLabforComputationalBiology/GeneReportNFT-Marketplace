package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Studio struct {
}

var (
	StudioController = &Studio{}
)

// CreateFromThirdParty 从第三方创建
func (s *Studio) CreateFromThirdParty(ctx *gin.Context) {
	var req dto.CreateAllFromThirdPartyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重试",
		})
		return
	}

	//fixme sse分阶段响应
	if toResp, err := services.StudioServ.CreateAllFromThirdPartyOnChain(ctx.GetString("user_address"), req); err != nil {
		_ = ctx.Error(err)
		return
	} else {
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    http.StatusOK,
			Message: "创建成功",
			Data:    toResp,
		})
	}

}

// GetProfileIds 获取后台中已经保存数据的该用户的profile id供用户选择
func (s *Studio) GetProfileIds(ctx *gin.Context) {
	userAddress := ctx.GetString("user_address")
	ids, err := services.StudioServ.GetProfileIdsByUser(userAddress)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    ids,
	})
}

//// GetWegeneTaskProgress 实时获取wegene请求任务进度
//func (s *Studio) GetWegeneTaskProgress(ctx *gin.Context) {
//	userAddress := ctx.GetString("user_address")
//}
