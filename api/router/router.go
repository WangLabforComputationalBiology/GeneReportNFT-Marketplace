package router

import (
	"GeneReport_platform/api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func registerSwaggerRouter(r *gin.Engine) {
	r.GET("/swagger/*any", func(c *gin.Context) {
		ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c)
	})

}

func registerUserRouter(r *gin.RouterGroup) {
	//用户主页信息
	r.GET("/info")
	//用户藏品信息
	r.GET("/collections")
	//用户订单信息
	r.GET("/orders")
}
func registerNFTRouter(r *gin.RouterGroup) {
	//藏品图片
	r.GET("/img")
	//藏品信息
	r.GET("/info")
	//生成订单
	r.POST("/buy")
}
func registerOrderRouter(r *gin.RouterGroup) {
	//订单信息
	r.GET("/info")
	//支付
	r.POST("/pay")
	//取消订单
	r.POST("/cancel")
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(otelgin.Middleware("GRM_Server"), middlewares.CORS())
	registerSwaggerRouter(r)
	User := r.Group("/:addr")
	registerUserRouter(User)
	NFT := r.Group("/:nft_id")
	registerNFTRouter(NFT)
	Order := r.Group("/:order_id")
	registerOrderRouter(Order)
	//商城首页
	r.GET("/")
	return r
}
