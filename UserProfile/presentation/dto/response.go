package dto

type AuthorizeStatus string

const (
	Authorized   AuthorizeStatus = "authorized"
	UnAuthorized AuthorizeStatus = "unauthorized"
	Error        AuthorizeStatus = "error"
)

type AuthorizeRespone struct {
	Status  AuthorizeStatus
	Message string
}
