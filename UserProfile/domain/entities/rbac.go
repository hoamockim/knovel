package entities

type Role struct {
	AutoIdEntity
	Name        string
	Description string
}

type Permission struct {
	AutoIdEntity
	Name     string
	SrvName  string // service name
	FuncName string // function name
}

type Rbac struct {
	AutoIdEntity
	RoleId       int
	PermissionID int
}

type UserRole struct {
	AutoIdEntity
	UserId string
	RoleId int
}
