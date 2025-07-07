package router

import (
	"GeneReport_platform/api/controllers"
	"GeneReport_platform/api/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
	"time"
)

func registerSwaggerRouter(r *gin.Engine) {
	//环境变量->条件启用swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func registerUserRouter(r *gin.RouterGroup) {

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
	r.GET("/info", middlewares.AuthMiddleware(), controllers.UserController.GetUserInfo)

	//获取用户头像
	r.GET("/profile", middlewares.AuthMiddleware(), controllers.UserController.GetProfileOfUser)
	//用户订单信息
	r.GET("/orders", middlewares.AuthMiddleware())
	//oauth2
	r.GET("/oauth2Wegene", controllers.UserController.Oauth2Wegene)
	//接收code
	r.GET("/receiveCode", controllers.UserController.ReceiveCode)
	//根据code获取用户的profile id供用户选择
	r.GET("/getProfileIds", middlewares.AuthMiddleware(), controllers.UserController.GetUsersProfileByCode)
	//用户授权哪份报告
	r.POST("/saveProfile", middlewares.AuthMiddleware(), controllers.UserController.SaveProfileInfo)
	//发送验证码
	r.POST("/send_email", middlewares.AuthMiddleware(), middlewares.ZapMiddleware(), controllers.UserController.SendEmailCode)
	//验证验证码
	r.POST("/verify_email", middlewares.AuthMiddleware(), middlewares.ZapMiddleware(), controllers.UserController.VerifyEmailCode)

}

func registerStudioRouter(r *gin.RouterGroup) {
	r.GET("/captcha", controllers.StudioController.GetCATCHA)
	r.POST("/captcha/check", controllers.StudioController.CheckCaptcha)

	r.GET("/getProfile/completed", middlewares.AuthMiddleware(), controllers.StudioController.GetCompletedProfileIds)
	r.GET("/getProfile/uncompleted", middlewares.AuthMiddleware(), controllers.StudioController.GetUncompletedProfileIDProgress)

	r.POST("/createFromThirdParty", middlewares.AuthMiddleware(), middlewares.ZapMiddleware(), controllers.StudioController.CreateFromThirdParty) //从第三方平台创建（链上操作）
}

func registerPlazaRouter(r *gin.RouterGroup) {
	r.GET("/:page", middlewares.ZapMiddleware(), controllers.MetadataController.GetAllMetadataOverview)

	//r.GET("/:page", middlewares.ZapMiddleware(), controllers.GeneSharingController.GetAllGeneSharingOverview)
}

func registerMetadataRouter(r *gin.RouterGroup) {
	r.GET("/:data_hash", middlewares.ZapMiddleware(), controllers.MetadataController.GetMetadataDetailByDataHash)
}

func registerGenoTypeRouter(r *gin.RouterGroup) {
	r.GET("/:data_hash", middlewares.AuthMiddleware(), middlewares.ZapMiddleware(), controllers.MetadataController.GetGenoTypeZip)

	r.POST("/obtainAccess", middlewares.AuthMiddleware(), middlewares.ZapMiddleware(), controllers.MetadataController.ObtainViewAccess)

}
func registerDownloadRouter(r *gin.RouterGroup) {
	r.GET("/:short_code", middlewares.ZapMiddleware(), middlewares.AuthMiddleware(), controllers.DownloadController.DownloadFile)
}
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORS())

	User := r.Group("/user")
	Studio := r.Group("/studio")
	Plaza := r.Group("/plaza")
	Metadata := r.Group("/metadata")
	GeneType := r.Group("/gene_type")
	Download := r.Group("/dl")
	registerSwaggerRouter(r)
	registerUserRouter(User)
	registerStudioRouter(Studio)
	registerMetadataRouter(Metadata)
	registerPlazaRouter(Plaza)
	registerGenoTypeRouter(GeneType)
	registerDownloadRouter(Download)

	r.GET("/events", func(c *gin.Context) {
		// 设置必要的SSE头信息
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 计数器
		count := 0

		// 持续发送消息
		for {
			select {
			case <-c.Request.Context().Done():
				// 客户端断开连接
				fmt.Println("Client disconnected")
				return
			default:
				count++
				// 格式化SSE消息
				message := fmt.Sprintf("data: Message %d at %v\n\n", count, time.Now())
				// 发送消息
				_, err := fmt.Fprint(c.Writer, message)
				if err != nil {
					fmt.Println("Write error:", err)
					return
				}
				// 刷新响应
				c.Writer.Flush()
				// 每秒发送一次
				time.Sleep(time.Second)
			}
		}
	})
	return r

}
