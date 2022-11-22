package model

import "time"
type Books struct{
	Name string `json:"name" db:"name"`
	Owner string `json:"owner" db:"owner"`
	Cost string `json:"cost" db:"cost"`
	ID string `json:"id" db:"id"`
	CreatedAt *time.Time `json:"createdat" db:"createdat"`
	UpdatedAt *time.Time `json:"updatedat" db:"updatedat"`
}
type GetAllBook struct{
	Book []Books `json:"book" db:"book"`
	Count int   `json:"count"`
}
type CreateBookRequest struct {
	Name string `json:"name" db:"name"`
	Owner string `json:"owner" db:"owner"`
	Cost string `json:"cost" db:"cost"`
}
type UbdateBookRequest struct {
	Name string `json:"name" db:"name"`
	Owner string `json:"owner" db:"owner"`
	Cost string `json:"cost" db:"cost"`
}
