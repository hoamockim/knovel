package middleware

import (
	"knovel/userprofile/presentation/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("X-Api-Key")
		if authorization != config.GetTaskClientKey() {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Next()
	}
}
