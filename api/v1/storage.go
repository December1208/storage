package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"mime/multipart"
	"storage/pkg"
)

type UploadFileData struct {
	File       *multipart.FileHeader
	PathPrefix string
}

type UploadFileReply struct {
	Identity string
	Url      string
}

type StorageServiceHTTPServer interface {
	UploadFile(context.Context, *UploadFileData) (*UploadFileReply, error)
}

type StorageService struct {
	server     StorageServiceHTTPServer
	router     gin.IRouter
	webContext *pkg.WebContext
}

func BindUploadFileData(webContext *pkg.WebContext) (UploadFileData, error) {
	var result UploadFileData
	file, err := webContext.FormFile("file")
	if err != nil {
		return result, err
	}
	result.File = file
	result.PathPrefix = webContext.Request.FormValue("path_prefix")
	return result, nil
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
		s.webContext.AbortWithError(err)
		return
	}
	out, err := s.server.(StorageServiceHTTPServer).UploadFile(newCtx, file)
	if err != nil {
		pkg.Logger.Error(err.Error())
		s.webContext.AbortWithError(err)
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
