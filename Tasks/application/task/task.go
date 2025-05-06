package task

import "context"

const (
	RS256 = "RS256"
)

type ListTaskRequest struct {
	PageSize int
	Offset   int
	UserId   string
}

type QueryTaskRequest struct {
	TaskId int
	UserId string
}

type CreateTaskRequest struct {
	Name        string `json:"name"`
	UserId      string `json:"user_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// update status of task
type UpdateTaskRequest struct {
	UserId string `json:"user_id"`
	TaskId int    `json:"task_id"`
	Status string `json:"status"`
}

type TaskResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type TaskManagement interface {
	CreateTask(context.Context, CreateTaskRequest) TaskResponse
	GetTasks(context.Context, ListTaskRequest) TaskResponse
	GetAssignedTasks(context.Context, QueryTaskRequest) TaskResponse
	UpdateTaskStatus(context.Context, UpdateTaskRequest) TaskResponse
	AssignTask(context.Context, UpdateTaskRequest) TaskResponse
	//DeleteTask(context.Context, DeleteTaskRequest) TaskResponse
}
