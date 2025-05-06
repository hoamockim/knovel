package application

import (
	"context"
	"knovel/tasks/application/task"
	"knovel/tasks/domain/entities"
	"knovel/tasks/domain/repositories"
	"net/http"
)

type Application struct {
	taskRepository repositories.TaskRepository
}

func NewApplication(taskRepository repositories.TaskRepository) *Application {
	return &Application{
		taskRepository: taskRepository,
	}
}

var _ task.TaskManagement = (*Application)(nil)

// CreateTask implements task.Authentication.
func (a *Application) CreateTask(ctx context.Context, request task.CreateTaskRequest) task.TaskResponse {
	if _, err := a.taskRepository.CreateTask(ctx, request.Name, request.UserId, request.Description, request.Status); err != nil {
		return task.TaskResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return task.TaskResponse{
		Code:    http.StatusCreated,
		Message: "create task successfully",
	}
}

// GetTasks implements task.Authentication.
func (a *Application) GetTasks(ctx context.Context, request task.ListTaskRequest) task.TaskResponse {
	tasks, err := a.taskRepository.GetTasks(ctx, request.PageSize, request.Offset)
	if err != nil {
		return task.TaskResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return task.TaskResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    tasks,
	}
}

func (a *Application) GetAssignedTasks(ctx context.Context, request task.QueryTaskRequest) task.TaskResponse {
	taskRecord, err := a.taskRepository.GetTasksByUserId(ctx, request.UserId)
	if err != nil {
		return task.TaskResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return task.TaskResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    taskRecord,
	}
}

// UpdateTask implements task.Authentication.
func (a *Application) UpdateTaskStatus(ctx context.Context, request task.UpdateTaskRequest) task.TaskResponse {
	if request.Status != string(entities.InProgress) && request.Status != string(entities.Completed) {
		return task.TaskResponse{
			Code:    http.StatusBadRequest,
			Message: "status is invalid",
		}
	}

	if _, err := a.taskRepository.UpdateTaskStatus(ctx, request.TaskId, request.UserId); err != nil {
		return task.TaskResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return task.TaskResponse{
		Code:    http.StatusOK,
		Message: "updated task successfully",
	}
}

// AssignTask implements task.TaskManagement.
func (a *Application) AssignTask(ctx context.Context, request task.UpdateTaskRequest) task.TaskResponse {
	if err := a.taskRepository.AssignTask(ctx, request.TaskId, request.UserId); err != nil {
		return task.TaskResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return task.TaskResponse{
		Code:    http.StatusOK,
		Message: "assigned task successfully",
	}
}
