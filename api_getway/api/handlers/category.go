package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)

func (h *Handler) CreateCategory(ctx *gin.Context) {
	var req book_shop.CategoryCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetCategoryService().CreateCategory(context.Background(), &req)

	if err != nil {
		h.log.Error(err.Error())
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetCategoryById(ctx *gin.Context) {

	var req book_shop.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetCategoryService().GetCategory(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}
