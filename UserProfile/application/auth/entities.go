package auth

type SignInRequest struct {
	Email    string `json:"email"`
	PassWord string `json:"pass_word"`
}

type SignOutRequest struct {
	Email string
}

type ClaimInfo struct {
	Id        string   `json:"id"`
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      []string `json:"roles"`
	Code      string   `json:"-"`
}

type AuthorizeInfo struct {
	UserId      string
	Roles       []string
	ServiceName string
	FuncName    string
	Permission  string
}

type ApplicationResponse struct {
	Status  int
	Message string
}

type PermissionOfService struct {
	FuncName   string `json:"func_name"`
	Permission string `json:"permission"`
}
