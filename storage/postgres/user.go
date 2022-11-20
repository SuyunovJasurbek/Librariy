package postgres

import (
	"library/model"

	"github.com/jmoiron/sqlx"
)

type userRepoImpl struct {
	db *sqlx.DB
}


func  (h userRepoImpl) GetAllUsers() ([]model.Users, error) {
	return nil, nil
}
func  (h userRepoImpl) CreateUser(entity model.CreateUserRequest, id string) error {
	return nil
}
func  (h userRepoImpl)  GetUserName(id string) (string, error) {
	return "", nil
}
func (h userRepoImpl)  UpdateUser(entity model.UbdateUserRequest, id string) (string, error) {
	return "", nil
}
func  (h userRepoImpl) DeleteUser(id string) error {
	return nil
}
