package router

import (
	"GeneReport_platform/api/controllers"
	"GeneReport_platform/api/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	_ "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
)

func registerSwaggerRouter(r *gin.Engine) {
	//环境变量->条件启用swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func testRouter(r *gin.RouterGroup) {
	//测试的接口
	r.GET("/user", controllers.UserController.Test)
	r.GET("/txinfo", controllers.GetTransactionInfo)
}

func registerUserRouter(r *gin.RouterGroup) {

	//todo get请求，  - url：/user/nonce/:账号 ,使用https，path和参数会被加密，account可以直接放在url
	//创建https服务器
	https := gin.Default()
	ssl := https.Group("/user")
	ssl.GET("/nonce/:user_address", controllers.UserController.GetNonce)
	go func() {
		log.Println("(((((((((( HTTPS server start at 8443 ))))))))))")
		// Start HTTPS server，server.crt：这是一个证书文件，server.key：这是一个私钥文件
		if err := https.RunTLS(":8443", "assets/TLS_files/server.crt", "assets/TLS_files/server.key"); err != nil {
			log.Fatalf("Failed to start HTTPS server: %v", err)
		}
	}()
	r.GET("/nonce/:user_address", controllers.UserController.GetNonce) //这个是http的，下面是https的
	//更改用户名
	r.POST("/edit/name", middlewares.AuthMiddleware(), controllers.UserController.EditUserName)
	//上传用户头像
	r.POST("/upload/avatar", middlewares.AuthMiddleware(), controllers.UserController.UploadAvatar)
	//登录
	r.POST("/login", controllers.UserController.Login)
	//登出
	r.POST("/logout", middlewares.AuthMiddleware(), controllers.UserController.Logout)
	//用户主页信息
	r.GET("/info", middlewares.AuthMiddleware(), controllers.UserController.GetInfo)
	//获取用户收藏列表
	r.GET("/gnfts", middlewares.AuthMiddleware(), controllers.UserController.GetGNFTList)
	//获取用户头像
	r.GET("/profile", middlewares.AuthMiddleware(), controllers.UserController.GetProfileOfUser)
	//用户订单信息
	r.GET("/orders", middlewares.AuthMiddleware())
	//oauth2
	r.GET("/oauth2Wegene", controllers.UserController.Oauth2Wegene)
	//接收code
	r.GET("/receiveCode", controllers.UserController.ReceiveCode)
	//根据code获取用户的profile id供用户选择
	r.GET("/getProfileIds" /*middlewares.AuthMiddleware(),*/, controllers.UserController.GetUsersProfileByCode)
	//用户授权哪份报告
	r.POST("/saveProfile" /*middlewares.AuthMiddleware(), */, controllers.UserController.SaveProfileInfo)
	//发送验证码
	r.POST("/send_sms", controllers.UserController.SendSMSCode)
	//验证验证码
	r.POST("/verify_sms", controllers.UserController.VerifySMSCode)
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
func registerStudioRouter(r *gin.RouterGroup) {
	r.GET("/captcha", controllers.StudioController.GetCatcha)
	r.POST("/captcha/check", controllers.StudioController.CheckCaptcha)
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(otelgin.Middleware("GRM_Server"), middlewares.CORS())

	//商城首页
	r.GET("/")

	Test := r.Group("/test")
	User := r.Group("/user")
	NFT := r.Group("/nft")
	Order := r.Group("/order")
	Studio := r.Group("/studio")
	registerSwaggerRouter(r)
	testRouter(Test)
	registerUserRouter(User)
	registerGNFTRouter(NFT)
	registerOrderRouter(Order)
	registerStudioRouter(Studio)

	//前缀是/test
	controllers.MyTestRoute(r)
	return r
}
