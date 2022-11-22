package handler

import (
	"library/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


func (h *HandlerImpl) GetName_book(ctx *gin.Context) {
	get_id := ctx.Param("id")
	get_data, err := h.strg.Book().GetBookName(get_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": model.Response{
				Data:    nil,
				Message: err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": model.Response{
			Data:    get_data,
			Message: "person found",
		},
	})

}

func (h *HandlerImpl) GetAllSearch_books(ctx *gin.Context) {
	offset, offset_exists := ctx.GetQuery("offset")
	if !offset_exists {
		offset = "0"
	}
	limit, limit_exists := ctx.GetQuery("limit")
	if !limit_exists {
		limit = "5"
	}
	search, search_exists := ctx.GetQuery("search")
	if !search_exists {
		search = ""
	}
	upd, err := h.strg.Book().GetAllSearchBooks(offset, limit, search)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": model.Response{
				Data:    err.Error(),
				Message: " ",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": model.Response{
			Data:    upd,
			Message: "filter data",
		},
	})
	
}
func (h *HandlerImpl) Create_book(ctx *gin.Context) {
	var book model.CreateBookRequest
	var book2 model.Books
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": model.Response{
				Data:    err.Error(),
				Message: "invalid request",
			},
		})
		return
	}
	id := uuid.New()
	timeNow := time.Now()
	book2.Name = book.Name
	book2.Owner = book.Owner
	book2.Cost = book.Cost
	book2.ID = id.String()
	book2.CreatedAt = &timeNow
	book2.UpdatedAt = &timeNow
	cret := h.strg.Book().CreateBook(book2, book2.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": model.Response{
				Data:    cret.Error(),
				Message: "invalid request",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": model.Response{
			Data:    book2.Name,
			Message: "create data",
		},
	})

}
func (h *HandlerImpl) Update_book(ctx *gin.Context) {
	var update_data model.UbdateBookRequest
	update_id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&update_data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": model.Response{
				Data:    err.Error(),
				Message: "invalid request",
			},
		})
		return
	}
	update_name, err := h.strg.Book().UpdateBook(update_data, update_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": model.Response{
				Data:    err.Error(),
				Message: "invalid request",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": model.Response{
			Data:    update_name,
			Message: "update data",
		},
	})

}
func (h *HandlerImpl) Delete_book(ctx *gin.Context) {
	delete_id := ctx.Param("id")
	err := h.strg.Book().DeleteBook(delete_id)
	if err != nil {
	//	fmt.Println(" handler  : ", delete_id)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": model.Response{
				Data:    err.Error(),
				Message: "invalid request ",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": model.Response{
			Data:    nil,
			Message: "delete data ",
		},
	})
}
