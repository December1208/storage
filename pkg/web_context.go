package pkg

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

const ServerErrorHTTPCode = 499

type WebContext struct {
	*gin.Context
}

func NewWebContext(c *gin.Context) *WebContext {
	webContext := WebContext{Context: c}

	return &webContext
}

func (ctx *WebContext) Success(obj interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    obj,
		"success": true,
		"detail":  "",
	})
}

func (ctx *WebContext) AbortWithError(err error) {
	code := -1
	status := ServerErrorHTTPCode
	msg := "未知错误"

	if err == nil {
		ctx.JSON(status, gin.H{
			"code":    code,
			"result":  nil,
			"message": msg,
			"success": false,
		})
		return
	}
	var c ICode
	if errors.As(err, &c) {
		status = c.HTTPCode()
		code = c.Code()
		msg = c.Message()
		return
	}

	ctx.JSON(status, gin.H{
		"code":    code,
		"result":  nil,
		"message": msg,
		"success": false,
	})
}

func (ctx *WebContext) BindJsonAndValidate(vData interface{}) error {
	if err := ctx.ShouldBindJSON(vData); err != nil {
		Logger.Warn("get param failed", zap.Error(err))
		return err
	}
	Logger.Info("get param success", zap.Any("data", vData))
	if err := Validate.Struct(vData); err != nil {
		Logger.Warn("validate failed", zap.Error(err))
		return err
	}
	return nil
}

func (ctx *WebContext) BindAndValidate(vData interface{}) error {
	if err := ctx.ShouldBind(vData); err != nil {
		Logger.Warn("get param failed", zap.Error(err))
		return err
	}
	Logger.Info("get param success", zap.Any("data", vData))
	if err := Validate.Struct(vData); err != nil {
		Logger.Warn("validate failed", zap.Error(err))
		return err
	}
	return nil
}

func (ctx *WebContext) BindQueryAndValidate(vData interface{}) error {
	if err := ctx.ShouldBindQuery(vData); err != nil {
		Logger.Warn("get param failed", zap.Error(err))
		return err
	}
	Logger.Info("get param success", zap.Any("data", vData))
	if err := Validate.Struct(vData); err != nil {
		Logger.Warn("validate failed", zap.Error(err))
		return err
	}
	return nil
}
