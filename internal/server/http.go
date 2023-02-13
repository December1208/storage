package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "storage/api/v1"
	"storage/internal/service"
	"time"
)

func NewHTTPServer(storageService *service.StorageService) *gin.Engine {
	eng := gin.New()

	eng.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		//定制日志格式
		return fmt.Sprintf("[%s] - [%s] [%s] [%s] %d %s %s\"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ErrorMessage,
		)
	}))
	//middleware.RegMiddleware(eng)
	v1.RegisterStorageHTTPServer(eng, storageService)
	return eng
}
