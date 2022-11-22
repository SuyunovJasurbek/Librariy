package main

import (
	"fmt"
	"library/config"
	"library/handler"
	"library/storage/postgres"

	"github.com/gin-gonic/gin"
)

const URL string = ":7070"

func main() {

	ctf := config.Load()
	fmt.Printf("%#+v\n", ctf)

	str := postgres.NewPostgres(fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
			ctf.PostgresUser,
			ctf.PostgresDatabase,
			ctf.PostgresPassword,
		),
	)
	h := handler.NewHandler(str)
	w := gin.Default()
	defer str.CloseDb()

	api_user := w.Group("api")
	{
		r := api_user.Group("/user")
		{

			r.GET("/", h.GetAll_user)
			//Get id
			r.GET("/:id", h.GetName_user)
			//Criate
			r.POST("/", h.Create_user)
			//Ubdate
			r.PUT(":id", h.Update_user)
			//Delete id
			r.DELETE("/:id", h.Delete_user)

		}
	}

	api_book := w.Group("api")
	{
		r := api_book.Group("/book")
		{
			//Get id
			r.GET("/:id", h.GetName_book)
			//Search books
			r.GET("/", h.GetAllSearch_books)
			//Criate
			r.POST("/", h.Create_book)
			//Ubdate
			r.PUT(":id", h.Update_book)
			//Delete id
			r.DELETE("/:id", h.Delete_book)
		}
	}

	w.Run(ctf.HTTPPort)
}
