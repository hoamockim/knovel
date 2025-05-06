package entities

type TaskStatus string

const (
	InProgress TaskStatus = "InProgress"
	Completed  TaskStatus = "Completed"
)

type Task struct {
	AutoIdEntity
	Name        string     `json:"name"`
	UserId      string     `json:"user_id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}
