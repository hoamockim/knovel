package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")

		accessToken := strings.Replace(authorization, "Bearer ", "", -1)

		token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return "", nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			id := fmt.Sprintf("%s", claims["id"])
			//roles := fmt.Sprintf("%s", claims["Roles"])
			ctx.Set("id", id)
			ctx.Set("roles", claims["roles"])
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
