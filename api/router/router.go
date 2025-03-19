package router

import (
	"GeneReport_platform/api/controllers"
	"GeneReport_platform/api/middlewares"
	_ "GeneReport_platform/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func registerSwaggerRouter(r *gin.Engine) {
	//环境变量->条件启用swagger
	//r.GET("/swagger/*any", func(c *gin.Context) {
	//	ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "SWAGGER")(c)
	//})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func registerUserRouter(r *gin.RouterGroup) {
	//获取nonce
	r.GET("/nonce/:user_address", controllers.UserController.GetNonce)

	//更改用户名
	r.POST("/edit/name", middlewares.AuthMiddleware(), controllers.UserController.EditUserName)
	//上传用户头像
	r.POST("/upload/avatar", middlewares.AuthMiddleware(), controllers.UserController.UploadProfile)
	//登录
	r.POST("/login", middlewares.AuthMiddleware(), controllers.UserController.Login)
	//登出
	r.POST("/logout", controllers.UserController.Logout)
	//用户主页信息
	r.GET("/info", controllers.UserController.GetInfo)
	//用户藏品信息
	r.GET("/gnfts", controllers.UserController.GetGNFTList)
	//我的--订单信息
	r.GET("/orders")
}
func registerGNFTRouter(r *gin.RouterGroup) {
	//藏品图片
	r.GET("/img")
	//藏品信息
	r.GET("/info")
	//生成订单
	r.POST("/buy", middlewares.AuthMiddleware())
}
func registerOrderRouter(r *gin.RouterGroup) {
	//订单信息
	r.GET("/info", middlewares.AuthMiddleware())
	//支付
	r.POST("/pay", middlewares.AuthMiddleware())
	//取消订单
	r.POST("/cancel", middlewares.AuthMiddleware())
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//启用https
	if err := r.RunTLS(":8443", "assets/TLS_files/server.crt", "assets/TLS_files/server.key"); err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
	r.Use(middlewares.CORS())
	User := r.Group("/user")
	NFT := r.Group("/nft_id")
	Order := r.Group("/order_id")

	registerSwaggerRouter(r)
	registerUserRouter(User)
	registerGNFTRouter(NFT)
	registerOrderRouter(Order)

	//商城首页
	r.GET("/")

	//前缀是/test
	controllers.MyTestRoute(r)
	return r
}
