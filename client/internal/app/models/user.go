package models

import "fmt"

type User struct {
	ID    int64
	Name  string
	Email string
	Role  int32
}

func (u *User) String() string {
	return fmt.Sprintf("id: %03v | name: %v | email: %v | role: %v", u.ID, u.Name, u.Email, u.Role)
}
