package models

import "fmt"

type User struct {
	ID    int64
	Name  string
	Email string
	Role  int32
}

func (u *User) String() string {
	return fmt.Sprintf("User (id = %v):\n  name: %v\n  email: %v\n  role: %v", u.ID, u.Name, u.Email, u.Role)
}
