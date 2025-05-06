package dto

import "knovel/userprofile/application/auth"

type SignInRequest struct {
	Email    string
	Password string
}

type SignInResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthorizeRequest struct {
	UserId      string   `json:"user_id"`
	Roles       []string `json:"roles"`
	ServiceName string   `json:"service_name"`
	FuncName    string   `json:"function_name"`
	Permission  string   `json:"permission"`
}

type PermissionOfServiceRequest struct {
	ServiceName string `json:"service_name"`
}

type PermissionOfServiceRespone struct {
	Data []*auth.PermissionOfService `json:"data"`
}
