package controllers

import "github.com/gin-gonic/gin"

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
