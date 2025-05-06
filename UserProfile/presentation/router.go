package presentation

import (
	"knovel/userprofile/presentation/handler"
	"knovel/userprofile/presentation/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	GetRouter() *gin.Engine
}

type GinRoute struct {
	engine      *gin.Engine
	authHandler handler.AuthHandler
	sysHandler  handler.SysHandler
}

func InitRouter(authHandler handler.AuthHandler, sysHandler handler.SysHandler) Router {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	router := &GinRoute{}
	router.engine = r
	router.authHandler = authHandler
	router.sysHandler = sysHandler
	apiV1 := router.engine.Group("api/v1")
	{
		apiV1.GET("/health-check", router.sysHandler.HealthCheck)
		apiV1.POST("/sign-in", router.authHandler.SignIn)
		apiV1.Use(middleware.BasicAuth())
		{
			apiV1.POST("/authorize", router.authHandler.Authorize)
			apiV1.GET("permissions", router.authHandler.FetchPermissionOfService)
		}
	}
	return router
}

func (router *GinRoute) GetRouter() *gin.Engine {
	return router.engine
}
