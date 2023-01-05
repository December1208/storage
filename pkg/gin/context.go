package gin

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"storage/pkg"
)

type WebContext struct {
	Metadata *pkg.Metadata
	Logger   *zap.Logger
	*gin.Context
}

func NewWebContext(c *gin.Context) *WebContext {
	baseContext := NewBaseContext()
	db, existed := c.Get("db")
	if existed && db != nil {
		baseContext.SetDB(db.(*gorm.DB))
	}
	ctxRedis, existed := c.Get("redis")
	if existed && ctxRedis != nil {
		baseContext.SetRedis(ctxRedis.(*redis.Client))
	}
	return &MyWebContext{baseContext, c, &util.SentryLog{L: baseContext.Log, Ctx: c}}
}

func (ctx *WebContext) Success(obj interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    obj,
		"success": true,
		"detail":  "",
	})
}

func (ctx *WebContext) AbortWithError(code, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"result":  nil,
		"success": false,
	})
}

func (ctx *WebContext) BindJsonAndValidate(vData interface{}) bool {
	if err := ctx.ShouldBindJSON(vData); err != nil {
		ctx.Logger.Warn("get param failed", zap.Error(err))
		ctx.AbortWithError("INVALID_PARAM", "参数错误")
		return false
	}
	ctx.Logger.Info("get param success", zap.Any("data", vData))
	if err := pkg.Validate.Struct(vData); err != nil {
		ctx.Logger.Warn("validate failed", zap.Error(err))
		ctx.AbortWithError("INVALID_PARAM", "参数错误")
		return false
	}
	return true
}

func (ctx *WebContext) BindAndValidate(vData interface{}) bool {
	if err := ctx.ShouldBind(vData); err != nil {
		ctx.Logger.Warn("get param failed", zap.Error(err))
		ctx.AbortWithError("INVALID_PARAM", "参数错误")
		return false
	}
	ctx.Logger.Info("get param success", zap.Any("data", vData))
	if err := pkg.Validate.Struct(vData); err != nil {
		ctx.Logger.Warn("validate failed", zap.Error(err))
		ctx.AbortWithError("INVALID_PARAM", "参数错误")
		return false
	}
	return true
}

func (ctx *WebContext) BindQueryAndValidate(vData interface{}) bool {
	if err := ctx.ShouldBindQuery(vData); err != nil {
		ctx.Logger.Warn("get param failed", zap.Error(err))
		ctx.AbortWithError("INVALID_PARAM", "参数错误")
		return false
	}
	ctx.Logger.Info("get param success", zap.Any("data", vData))
	if err := pkg.Validate.Struct(vData); err != nil {
		ctx.Logger.Warn("validate failed", zap.Error(err))
		ctx.AbortWithError("INVALID_PARAM", "参数错误")
		return false
	}
	return true
}
