package model

import "time"

type Users struct {
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Age       string       `json:"age" db:"age"`
	ID string 			  `json:"id" db:"id"`
	CreatedAt *time.Time `json:"createdat" db:"createdat"`
	UpdatedAt *time.Time `json:"updatedat" db:"updatedat"`
}
type GetAllUser struct{
	User []Users `json:"user" db:"user"`
	Count int   `json:"count"`
}
type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Age      string   `json:"age" db:"age"`
}

type UbdateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email string `json:"email" db:"email"`
	Age string `json:"age" db:"age"`
}
