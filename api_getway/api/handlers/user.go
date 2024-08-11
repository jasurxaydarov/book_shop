package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)

func (h *Handler) CreateUser(ctx *gin.Context) {

	var req book_shop.UserCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetUserSevice().CreateUser(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetUserById(ctx *gin.Context) {

	var req book_shop.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetUserSevice().GetUser(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}
