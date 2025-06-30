package middlewares

import (
	"GeneReport_platform/pkg/appErrors"
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var (
	AppLogger *zap.Logger

	encoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,        //默认换行符
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 日志级别大写并带颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder,       // 时间格式如 "2025-05-07T10:00:00Z"
		EncodeDuration: zapcore.StringDurationEncoder,    //耗时
		EncodeCaller:   zapcore.FullCallerEncoder,        // 调用者信息
		EncodeName:     zapcore.FullNameEncoder,
	}
)

func init() {
	// 创建 JSON 编码器
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// 创建写入器（控制台输出）
	writeSyncer := zapcore.AddSync(zapcore.Lock(os.Stdout))

	// 创建日志核心
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	// 创建 Logger
	AppLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func ZapMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		method := ctx.Request.Method

		// 读取请求体
		var requestBody string
		if ctx.Request.Body != nil {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				// 恢复请求体以供后续处理
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// 处理请求
		ctx.Next()

		var appErr appErrors.AppError
		//从gin上下文中响应错误
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			appErr, _ = err.(appErrors.AppError)
			if appErr.Data != nil {
				ctx.JSON(appErr.Code, gin.H{
					"Code":    appErr.Code,
					"Message": appErr.Message,
					"Data":    appErr.Data,
				})
				return
			}
			ctx.JSON(appErr.Code, appErr.ToErrResponse())
		}

		// 计算请求耗时
		latency := time.Since(start)
		status := ctx.Writer.Status()

		// 如果状态码表示错误（>=500），记录错误日志
		if status >= 500 {
			AppLogger.Error("Internal Server Error",
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Int("status", status),
				zap.Duration("latency", latency),
				zap.String("client_ip", ctx.ClientIP()),
				zap.String("request_body", requestBody),        // 添加请求体
				zap.String("errors", appErr.ErrorWithDetail()), // 记录 Gin 上下文中的错误
			)
		} else {
			AppLogger.Info("Info output",
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.Int("status", status),
				zap.Duration("latency", latency),
				zap.String("client_ip", ctx.ClientIP()),
				zap.String("request_body", requestBody),        // 添加请求体
				zap.String("errors", appErr.ErrorWithDetail()), // 记录 Gin 上下文中的错误
			)
		}
	}
}
