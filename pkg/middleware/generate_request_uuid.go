package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestUUIDKey = "X-REQUEST-UUID"

func GenerateRequestUUIDMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestUUID := context.GetHeader(RequestUUIDKey)
		if requestUUID == "" {
			requestUUID = uuid.NewString()
			context.Header(RequestUUIDKey, requestUUID)
		}
	}
}
