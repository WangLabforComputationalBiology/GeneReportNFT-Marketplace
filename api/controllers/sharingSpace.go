package controllers

import (
	"github.com/gin-gonic/gin"
)

type SharingSpace struct{}

var (
	SharingSpaceController = &SharingSpace{}
)

func (s *SharingSpace) GetListings(ctx *gin.Context) {

}
