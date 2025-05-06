package handler

import (
	"context"
	"errors"
	"fmt"
	"knovel/tasks/application"
	"knovel/tasks/application/task"
	"knovel/tasks/presentation/client"
	"knovel/tasks/presentation/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FuncName string

const (
	ServiceName               = "task"
	CreateTask       FuncName = "create_task"
	UpdateTaskStatus FuncName = "update_task_status"
	AssignTask       FuncName = "assign_task"
	ViewTask         FuncName = "view_task"
)

type TaskHandler interface {
	CreateTask(ctx *gin.Context)
	AssignTask(ctx *gin.Context)
	ReadTasksAssigned(ctx *gin.Context)
	ListTask(ctx *gin.Context)
	UpdateTaskStatus(ctx *gin.Context)
}

type TaskController struct {
	task       task.TaskManagement
	authClient client.RestClient
	permission map[FuncName]string //funcName, permission
}

var _ TaskHandler = (*TaskController)(nil)

func NewTaskHandler(task *application.Application, authClient client.RestClient) TaskHandler {
	taskController := &TaskController{
		task:       task,
		authClient: authClient,
	}

	taskController.loadPermissionsOfservice()
	return taskController
}

func parseResponse(ctx *gin.Context, result task.TaskResponse) {
	ctx.JSON(result.Code, dto.TaskResponse{
		Message: result.Message,
		Data:    result.Data,
	})
}

func (handler *TaskController) CreateTask(ctx *gin.Context) {
	userId, ok := ctx.Get("id")
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//-authorize
	roles, err := parseCtxForAuthorizeRequest(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// handler.authorize(userId.(string), roles, serviceName, functionName, permissionName)
	permissionName := handler.permission[CreateTask]
	authorStatus := handler.authorize(userId.(string), roles, ServiceName, CreateTask, permissionName)
	if authorStatus == dto.UnAuthorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if authorStatus == dto.Error {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//-end authorize

	// Binding request body to variable
	var request dto.TaskRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Insert task to database
	result := handler.task.CreateTask(ctx.Request.Context(), task.CreateTaskRequest{
		Name:        request.Name,
		UserId:      userId.(string),
		Description: request.Description,
		Status:      request.Status,
	})

	parseResponse(ctx, result)

}

// AssignTask implements TaskHandler.
func (handler *TaskController) AssignTask(ctx *gin.Context) {
	userId, ok := ctx.Get("id")
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//authorize
	roles, err := parseCtxForAuthorizeRequest(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	permissionName := handler.permission[AssignTask]
	authorStatus := handler.authorize(userId.(string), roles, ServiceName, AssignTask, permissionName)
	if authorStatus == dto.UnAuthorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if authorStatus == dto.Error {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//end authorize
	// Binding request body to variable
	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var request dto.AssignTaskRequest
	ctx.BindJSON(&request)

	// Update status of task in database

	result := handler.task.AssignTask(ctx.Request.Context(), task.UpdateTaskRequest{
		UserId: request.UserId,
		TaskId: taskId,
	})

	parseResponse(ctx, result)
}

func (handler *TaskController) ReadTasksAssigned(ctx *gin.Context) {
	userId, ok := ctx.Get("id")
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//authorize
	roles, err := parseCtxForAuthorizeRequest(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	permissionName := handler.permission[ViewTask]
	authorStatus := handler.authorize(userId.(string), roles, ServiceName, ViewTask, permissionName)
	if authorStatus == dto.UnAuthorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if authorStatus == dto.Error {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//end authorize
	// Get task id from params

	// Fetch data
	result := handler.task.GetAssignedTasks(ctx.Request.Context(), task.QueryTaskRequest{
		UserId: userId.(string),
	})

	parseResponse(ctx, result)

}

func (handler *TaskController) ListTask(ctx *gin.Context) {
	// Validate jwt token
	userId, ok := ctx.Get("id")
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//authorize
	roles, err := parseCtxForAuthorizeRequest(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	permissionName := handler.permission[ViewTask]
	authorStatus := handler.authorize(userId.(string), roles, ServiceName, ViewTask, permissionName)
	if authorStatus == dto.UnAuthorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if authorStatus == dto.Error {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//end authorize

	// Get page number and page size
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	pageSize, err := strconv.Atoi(ctx.Query("pagesize"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	// Fetch data
	result := handler.task.GetTasks(ctx.Request.Context(), task.ListTaskRequest{
		PageSize: pageSize,
		Offset:   (page - 1) * pageSize,
		UserId:   userId.(string),
	})

	parseResponse(ctx, result)
}

func (handler *TaskController) UpdateTaskStatus(ctx *gin.Context) {
	// Validate jwt token
	userId, ok := ctx.Get("id")
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// authorize
	roles, err := parseCtxForAuthorizeRequest(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	permissionName := handler.permission[UpdateTaskStatus]
	authorStatus := handler.authorize(userId.(string), roles, ServiceName, UpdateTaskStatus, permissionName)
	if authorStatus == dto.UnAuthorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if authorStatus == dto.Error {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	//end authorize

	// Get task id from params
	taskId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var request dto.TaskRequest
	ctx.BindJSON(&request)

	// Update status of task in database

	result := handler.task.UpdateTaskStatus(ctx.Request.Context(), task.UpdateTaskRequest{
		UserId: userId.(string),
		TaskId: taskId,
		Status: request.Status,
	})

	parseResponse(ctx, result)
}

func parseCtxForAuthorizeRequest(ctx *gin.Context) ([]string, error) {

	r, ok := ctx.Get("roles")
	if !ok {
		return nil, errors.New("invalid roles")
	}
	roles := make([]string, 0)
	if arr, ok := r.([]interface{}); ok {
		for _, item := range arr {
			roles = append(roles, item.(string))
		}
	} else {
		return nil, errors.New("invalid roles")
	}

	return roles, nil
}

func (handler *TaskController) authorize(userId string, roles []string, serviceName string, funcName FuncName, permission string) dto.AuthorizeStatus {
	authorEndpoint := handler.authClient.MakeEndpoint("authorize", http.MethodPost)

	res, err := authorEndpoint(context.Background(), dto.AuthorizeRequest{
		UserId:      userId,
		Roles:       roles,
		Permission:  permission,
		ServiceName: serviceName,
		FuncName:    string(funcName),
	})
	if err != nil {
		return dto.Error
	}
	if authorRes, ok := res.(*dto.AuthorizeRespone); ok {
		return authorRes.Status
	}
	return dto.Error
}

func (handler *TaskController) loadPermissionsOfservice() {
	handler.permission = make(map[FuncName]string)
	endpoint := handler.authClient.MakeEndpoint(fmt.Sprintf("permissions?service_name=%s", ServiceName), http.MethodGet)
	res, _ := endpoint(context.Background(), nil)
	if permissions, ok := res.(*dto.PermissionOfServiceRespone); ok {
		for _, permission := range permissions.Data {
			handler.permission[FuncName(permission.FuncName)] = permission.Permission
		}
	}
}
