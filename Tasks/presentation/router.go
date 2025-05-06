package presentation

import (
	"knovel/tasks/presentation/handler"
	"knovel/tasks/presentation/middleware"

	"github.com/gin-gonic/gin"
)

type Router interface {
	GetRouter() *gin.Engine
}

type GinRouter struct {
	engine  *gin.Engine
	handler handler.TaskHandler
}

func InitRouter(h handler.TaskHandler) Router {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	router := GinRouter{}
	router.engine = r
	router.handler = h
	apiV1 := router.engine.Group("api/v1")

	apiV1.Use(middleware.Jwt())
	{
		apiV1.POST("/task", router.handler.CreateTask)
		apiV1.GET("/tasks", router.handler.ListTask)
		apiV1.GET("/tasks-assigned", router.handler.ReadTasksAssigned)
		apiV1.PATCH("/task/status/:id", router.handler.UpdateTaskStatus)
		apiV1.PATCH("/task/assign/:id", router.handler.AssignTask)
	}
	return &router
}

func (router *GinRouter) GetRouter() *gin.Engine {
	return router.engine
}
