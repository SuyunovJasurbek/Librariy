package model

import "time"

type Users struct {
	Username  string     `json:"username" db:"createdat"`
	Email     string     `json:"email" db:"createdat"`
	Age       int8       `json:"age" db:"createdat"`
	CreatedAt *time.Time `json:"createdat" db:"createdat"`
	UpdatedAt *time.Time `json:"updatedat" db:"updatedat"`
}
type CreateUserRequest struct {
	Username string `json:"username" db:"createdat"`
	Email    string `json:"email" db:"createdat"`
	Age      int8   `json:"age" db:"createdat"`
	ID string `json:"id" db:"id"`
}
type UbdateUserRequest struct {
	Name string `json:"name" db:"name"`
	Owner string `json:"owner" db:"owner"`
	Cost string `json:"cost" db:"id"`
}
