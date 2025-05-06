package dto

type AuthorizeStatus string

const (
	Authorized   AuthorizeStatus = "authorized"
	UnAuthorized AuthorizeStatus = "unauthorized"
	Error        AuthorizeStatus = "error"
)

type TaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type AssignTaskRequest struct {
	UserId string `json:"user_id"`
}

type TaskResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type AuthorizeRequest struct {
	UserId      string   `json:"user_id"`
	Roles       []string `json:"roles"`
	Permission  string   `json:"permission"`
	ServiceName string   `json:"service_name"`
	FuncName    string   `json:"function_name"`
}

type AuthorizeRespone struct {
	Status  AuthorizeStatus
	Message string
}

type PermissionOfServiceRespone struct {
	Data []*struct {
		FuncName   string `json:"func_name"`
		Permission string `json:"permission"`
	} `json:"data"`
}
