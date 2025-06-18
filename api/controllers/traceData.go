package controllers

import (
	"github.com/gin-gonic/gin"
)

type TraceData struct{}

var (
	TraceDataController = &TraceData{}
)

func (t *TraceData) GetTraceDataByMetadataID(ctx *gin.Context) {

}
