package router

import (
	"GeneReport_platform/api/controllers"
	"GeneReport_platform/api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
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

	//测试的接口
	r.GET("/test", controllers.UserController.Test)

	r.GET("/nonce/:user_address", controllers.UserController.GetNonce) //这个是http的，下面是https的

	//todo get请求，  - url：/user/nonce/:账号 ,使用https，path和参数会被加密，account可以直接放在url
	//创建https服务器
	https := gin.Default()
	ssl := https.Group("/user")
	ssl.GET("/nonce/:user_address", controllers.UserController.GetNonce)
	go func() {
		// Start HTTPS server，server.crt：这是一个证书文件，server.key：这是一个私钥文件
		if err := https.RunTLS(":8443", "assets/TLS_files/server.crt", "assets/TLS_files/server.key"); err != nil {
			log.Fatalf("Failed to start HTTPS server: %v", err)
		}
	}()

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
	//获取用户头像
	r.GET("/profile", controllers.UserController.GetProfileOfUser)
	//用户订单信息
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
	r.Use(otelgin.Middleware("GRM_Server"), middlewares.CORS())
	registerSwaggerRouter(r)
	User := r.Group("/user")
	registerUserRouter(User)
	NFT := r.Group("/nft_id")
	registerNFTRouter(NFT)
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
