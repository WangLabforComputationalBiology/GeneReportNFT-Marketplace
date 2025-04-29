package middlewares

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"runtime"
	"strconv"
)

var (
	EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding, // 每条日志换行
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	ZapConfig = zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.ErrorLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    EncoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
)

// ErrorLog 统一错误日志格式
type ErrorLog struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`          // 错误详情
	Stack   string `json:"stack,omitempty"` // 仅在调试模式下返回
}

// AppError 自定义错误类型
type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  error  `json:"detail"`
}

// AppError 隐式实现 error 接口
func (e *AppError) Error() string {
	return e.Message
}

// ErrorHandlerConfig 中间件配置
type ErrorHandlerConfig struct {
	DebugMode   bool  // 是否返回堆栈信息
	LogStack    bool  // 是否记录堆栈信息
	MaxBodySize int64 // 最大请求Body读取大小
}

// ErrorHandler 增强的错误处理中间件，使用自定义 Zap 编码器输出到控制台
func ErrorHandler(config ErrorHandlerConfig) gin.HandlerFunc {

	// 构建 Zap 日志器
	logger, err := ZapConfig.Build()
	if err != nil {
		fmt.Println("Failed to initialize logger:", err)
		logger, _ = zap.NewDevelopment() // 回退到默认
	}
	// 使用defer关键字，在函数返回之前执行该函数
	defer func(logger *zap.Logger) {
		// 尝试同步logger
		if err := logger.Sync(); err != nil {
			// 如果同步失败，打印错误信息
			fmt.Println("Failed to sync logger:", err)
		}
	}(logger)

	if config.MaxBodySize == 0 {
		config.MaxBodySize = 1024 * 1024 // 默认1MB
	}

	return func(c *gin.Context) {
		// 生成请求ID
		requestID := uuid.New().String()
		c.Set("request_id", requestID)
		logger := logger.With(zap.String("request_id", requestID))

		// 执行后续处理程序
		defer func() {
			// 捕获 panic
			if cause := recover(); cause != nil {
				stack := ""
				if config.LogStack || config.DebugMode {
					stack = formatStack(3)
				}
				// 使用 Zap 输出到控制台
				logger.Error("Panic recovered",
					zap.Any("error", cause),
					zap.String("stack", stack),
					zap.String("request", dumpRequest(c.Request, config.MaxBodySize)),
				)

				response := ErrorResponse{
					Code:    "INTERNAL_SERVER_ERROR",
					Message: "Internal Server Error",
				}
				if config.DebugMode {
					response.Stack = stack
				}
				c.JSON(http.StatusInternalServerError, response)
				c.Abort()
				return
			}

			// 检查 Gin 上下文错误
			if len(c.Errors) > 0 {
				err := c.Errors.Last()
				stack := ""
				if config.LogStack || config.DebugMode {
					stack = formatStack(3)
				}

				// 确定状态码和响应
				status := http.StatusInternalServerError
				response := ErrorResponse{
					Code:    "UNKNOWN_ERROR",
					Message: err.Error(),
				}

				// 处理自定义错误
				if customErr, ok := err.Err.(*AppError); ok {
					response.Code = customErr.Code
					response.Message = customErr.Message
					switch customErr.Code {
					case "BAD_REQUEST":
						status = http.StatusBadRequest
					case "NOT_FOUND":
						status = http.StatusNotFound
					case "UNAUTHORIZED":
						status = http.StatusUnauthorized
					case "FORBIDDEN":
						status = http.StatusForbidden
					case "CONFLICT":
						status = http.StatusConflict
					case "UNSUPPORTED_MEDIA_TYPE":
						status = http.StatusUnsupportedMediaType
					case "TIMEOUT":
						status = http.StatusRequestTimeout
					case "EXTERNAL_SERVICE_ERROR":
						status = http.StatusServiceUnavailable
					case "DATABASE_ERROR":
						status = http.StatusInternalServerError
					default:
						status = http.StatusInternalServerError
					}
				} else if err.Type == gin.ErrorTypeBind {
					response.Code = "BAD_REQUEST"
					response.Message = "Invalid request parameters"
					status = http.StatusBadRequest
				}

				// 使用 Zap 输出到控制台
				logger.Error("Request error",
					zap.String("error", err.Error()),
					zap.String("type", strconv.FormatUint(uint64(err.Type), 10)),
					zap.String("stack", stack),
					zap.String("request", dumpRequest(c.Request, config.MaxBodySize)),
				)

				// 设置堆栈（调试模式）
				if config.DebugMode {
					response.Stack = stack
				}

				c.JSON(status, response)
				c.Abort()
				return
			}
		}()

		// 继续处理请求
		c.Next()
	}
}

// NewCustomError 创建自定义错误
func NewCustomError(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Detail:  err,
	}
}

// formatStack 格式化堆栈信息
func formatStack(skip int) string {
	var buf bytes.Buffer
	pcs := make([]uintptr, 32)
	n := runtime.Callers(skip, pcs)
	for i := 0; i < n; i++ {
		fn := runtime.FuncForPC(pcs[i])
		if fn == nil {
			continue
		}
		file, line := fn.FileLine(pcs[i])
		fmt.Fprintf(&buf, "%s\n\t%s:%d\n", function(pcs[i]), file, line)
	}
	return buf.String()
}

// function 返回简化的函数名
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return []byte("???")
	}
	name := []byte(fn.Name())
	if lastslash := bytes.LastIndex(name, []byte("/")); lastslash >= 0 {
		name = name[lastslash+1:]
	}
	if period := bytes.Index(name, []byte(".")); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, []byte("·"), []byte("."), -1)
	return name
}

// dumpRequest 格式化请求信息
func dumpRequest(req *http.Request, maxBodySize int64) string {
	var b bytes.Buffer
	reqURI := req.RequestURI
	if reqURI == "" {
		reqURI = req.URL.RequestURI()
	}
	fmt.Fprintf(&b, "%s %s HTTP/%d.%d\n", req.Method, reqURI, req.ProtoMajor, req.ProtoMinor)
	if req.Body != nil {
		n, err := io.Copy(&b, io.LimitReader(req.Body, maxBodySize))
		if err != nil {
			return fmt.Sprintf("Error reading body: %v", err)
		}
		if n > 0 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

func errResponse(ctx *gin.Context) {
	ctx.Error()
}
