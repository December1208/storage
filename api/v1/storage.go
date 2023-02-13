package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"mime/multipart"
	"storage/pkg"
)

type UploadFileRequest struct {
	File       *multipart.FileHeader
	PathPrefix string
}

type UploadFileReply struct {
	Identity string
	Url      string
}

type StorageServiceHTTPServer interface {
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileReply, error)
}

func BindUploadFileData(webContext *pkg.WebContext) (UploadFileRequest, error) {
	var result UploadFileRequest
	file, err := webContext.FormFile("file")
	if err != nil {
		return result, err
	}
	result.File = file
	result.PathPrefix = webContext.Request.FormValue("path_prefix")
	return result, nil
}

func UploadFileHandler(srv StorageServiceHTTPServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		webContext := pkg.NewWebContext(ctx)
		md := metadata.New(nil)
		for k, v := range ctx.Request.Header {
			md.Set(k, v...)
		}
		newCtx := metadata.NewIncomingContext(webContext, md)
		reqData, err := BindUploadFileData(webContext)
		if err != nil {
			pkg.Logger.Error(err.Error())
			webContext.AbortWithError(err)
			return
		}
		out, err := srv.UploadFile(newCtx, &reqData)
		if err != nil {
			pkg.Logger.Error(err.Error())
			webContext.AbortWithError(err)
			return
		}
		webContext.Success(out)
	}

}

func RegisterStorageHTTPServer(eng *gin.Engine, srv StorageServiceHTTPServer) {

	v2api := eng.Group("/v1/api")

	v2api.POST("/files/normal_upload", UploadFileHandler(srv))
}
