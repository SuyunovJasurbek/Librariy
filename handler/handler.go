package handler

import "library/storage"

type HandlerImpl struct{
	strg storage.StorageI
}
func NewHandler(s storage.StorageI ) HandlerImpl{
	return HandlerImpl{
		strg: s,
	}
}