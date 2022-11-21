package storage

import "library/model"

type StorageI interface{ 
	User() UserRepoI
	Book() BookRepoI
	CloseDb() error
}

type BookRepoI interface{
	 CreateBook(entity model.Books, id string)error
	 GetBookName(id string) (string , error) 
	 GetAllSearchBooks(offset, limit, search string) (*model.GetAllBook, error)
	 UpdateBook(entity model.UbdateBookRequest, id string) (*model.Books, error) 
	 DeleteBook(id string) error
}

type UserRepoI interface{
	GetAllSearchUser(offset, limit, search string) (*model.GetAllUser, error)
	CreateUser(entity model.Users, id string)error
	GetUserName(id string) (string , error) 
	UpdateUser(entity model.UbdateUserRequest, id string) (*model.Users, error) 
	DeleteUser(id string) error
}