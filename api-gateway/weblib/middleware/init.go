package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(serviceMap map[string]interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 将实例存在gin.Keys中
		context.Keys = make(map[string]interface{})
		context.Keys = serviceMap
		context.Next()
	}
}
