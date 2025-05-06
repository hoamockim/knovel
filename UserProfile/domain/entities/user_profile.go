package entities

type User struct {
	Entity
	Email     string
	UserName  string
	Password  string
	FirstName string
	LastName  string
}
