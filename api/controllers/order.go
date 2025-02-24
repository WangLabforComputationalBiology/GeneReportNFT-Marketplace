package controllers

import "github.com/gin-gonic/gin"

/*
这段代码定义了一个 OrderController 结构体，并为其实现了几个方法，这些方法将用于处理与订单相关的操作
*/
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
