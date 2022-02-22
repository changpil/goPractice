package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

//Slice of Pointing users
var (
	users  []*User
	nextID = 1
)
