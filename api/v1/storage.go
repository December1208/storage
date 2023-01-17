package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"mime/multipart"
	"storage/api"
	"storage/pkg"
)

type StorageServiceHTTPServer interface {
	UploadFile(context.Context, *multipart.FileHeader) (*HelloReply, error)
}

type StorageService struct {
	server     StorageServiceHTTPServer
	router     gin.IRouter
	webContext *pkg.WebContext
}

func (s *StorageService) UploadFile(ctx *gin.Context) {
	webContext := pkg.NewWebContext(ctx)
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(webContext, md)
	file, err := webContext.FormFile("file")
	if err != nil {
		pkg.Logger.Error(err.Error())
		s.webContext.AbortWithError(api.NOT_DEFINED, "请求参数错误")
		return
	}
	out, err := s.server.(StorageServiceHTTPServer).UploadFile(newCtx, file)
	if err != nil {
		pkg.Logger.Error(err.Error())
		s.webContext.AbortWithError(api.NOT_DEFINED, "服务错误")
		return
	}
	s.webContext.Success(out)
}

func RegisterStorageHTTPServer(eng *gin.Engine) {

	//healthController := new(demo.HealthController)
	//eng.GET("/health", healthController.Health)
	//
	//v2api := eng.Group("/v2/api")
	//{
	//	sse_server.AddApiRouter(v2api)
	//}
	//v2InnerApi := eng.Group("/v2/inner_api")
	//{
	//	sse_server.AddInnerApiRouter(v2InnerApi)
	//}
}
