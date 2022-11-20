package storage

import "library/model"

type StorageI interface{ 
	User() UserRepoI
	Book() BookRepoI
	CloseDb() error
}

type BookRepoI interface{
	 GetAllBooks() ([]model.Books, error) 
	 CreateBook(entity model.Books, id string)error
	 GetBookName(id string) (string , error) 
	 GetAllSearchBooks(offset, limit, search string) (*model.GetAllBook, error)
	 UpdateBook(entity model.UbdateBookRequest, id string) (string, error) 
	 DeleteBook(id string) error
}

type UserRepoI interface{
	GetAllUsers() ([]model.Users, error) 
	CreateUser(entity model.CreateUserRequest, id string)error
	GetUserName(id string) (string , error) 
	UpdateUser(entity model.UbdateUserRequest, id string) (string, error) 
	DeleteUser(id string) error
}