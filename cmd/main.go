package main

import (
	"library/handler"
	"library/storage/postgres"
	"github.com/gin-gonic/gin"
)

func main() {

	str :=postgres.NewPostgres()
	h:=handler.NewHandler(str)
	const URL string = ":7070"
	w := gin.Default()
	api_user := w.Group("api")
	{
		r:= api_user.Group("/user")
		{
			//Get all Users 
			r.GET("/",h.GetAll_user)
			//Get id
			r.GET("/:id",h.GetName_user)
			 //Criate
			r.POST("/",h.Create_user)
			//Ubdate
			r.PUT(":id",h.Update_user)
			//Delete id
			r.DELETE("/:id",h.Delete_user)
			
		}
	}
	api_book := w.Group("api")
	{
		r:= api_book.Group("/book")
		{
			//Get all Users 
			r.GET("/",h.GetAll_book)
			//Get id
			r.GET("/:id",h.GetName_book)
			//Search books
			r.GET("/b",h.GetAllSearch_books)
			 //Criate
			r.POST("/",h.Create_book)
			//Ubdate
			r.PUT(":id",h.Update_book)
			//Delete id
			r.DELETE("/:id",h.Delete_book)
		}
	}

	w.Run(URL)
}