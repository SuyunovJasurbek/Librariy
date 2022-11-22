package handler

import (
	"fmt"
	"library/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
func (h *HandlerImpl) GetAll_user(ctx *gin.Context){

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

func (h *HandlerImpl) GetName_user(ctx *gin.Context) {
	get_id := ctx.Param("id")
	get_name, err := h.strg.User().GetUserName(get_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message": model.Response{
				Data:    err.Error(),
				Message: "not faund ",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": model.Response{
			Data:    get_name,
			Message: "sucses",
		},
	})

}

func (h *HandlerImpl) Create_user(ctx *gin.Context) {
	var user model.CreateUserRequest
	var cre_user model.Users
	err := ctx.ShouldBindJSON(&user)
	fmt.Println(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": model.Response{
				Data:    err.Error(),
				Message: "invalid request",
			},
		})
		return
	}
	id := uuid.New()
	tim :=time.Now()
	cre_user.Username=user.Username
	cre_user.Email=user.Email
	cre_user.Age=user.Age
	cre_user.CreatedAt=&tim
	cre_user.UpdatedAt=&tim

	fmt.Println(cre_user)
	err2 := h.strg.User().CreateUser(cre_user, id.String())
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": model.Response{
				Data:    err2,
				Message: "invalid request",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": model.Response{
			Data:    cre_user.Username,
			Message: "sucses",
		},
	})

}
func (h *HandlerImpl) Update_user(ctx *gin.Context) {

	var upd_user model.UbdateUserRequest
	upd_id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&upd_user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": model.Response{
				Data:    err.Error(),
				Message: "invalit request",
			},
		})
		return
	}

	
	upd_data, err := h.strg.User().UpdateUser(upd_user, upd_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": model.Response{
				Data:    err.Error(),
				Message: "not faunt",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": model.Response{
			Data:    upd_data,
			Message: "sucses",
		},
	})
}

func (h *HandlerImpl) Delete_user(ctx *gin.Context) {
	del_id :=ctx.Param("id")

	err:=h.strg.User().DeleteUser(del_id)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Message": model.Response{
				Data:    err.Error(),
				Message: "invalit request",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": model.Response{
			Data:    nil,
			Message: "delete ",
		},
	})

}
