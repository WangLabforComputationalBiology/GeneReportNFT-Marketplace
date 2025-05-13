package controllers

import (
	"GeneReport_platform/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

var OrderCtrller = &OrderController{}

func (cOrder *OrderController) GetInfo(ctx *gin.Context) {
}

func (cOrder *OrderController) GetList(ctx *gin.Context) {

}

func (cOrder *OrderController) Create(ctx *gin.Context) {

}

func (cOrder *OrderController) Cancel(ctx *gin.Context) {

}

// Pay 支付
func (cOrder *OrderController) Pay(ctx *gin.Context) {

}

// 获取基于的数据,请求是router.GET("/orders/:param", OrderCtrller.GetInfo)
func (cOrder *OrderController) GetData(ctx *gin.Context) {
	services.GetDataImpl(ctx)
}
