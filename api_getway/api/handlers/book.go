package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *Handler) CreateBook(ctx *gin.Context) {

	var req book_shop.BookCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetBookService().CreateBook(context.Background(), &req)

	if err != nil {
		h.log.Error("err on  GetBookService().CreateBook", logger.Error(err))
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)
}

func (h *Handler) GetBookById(ctx *gin.Context) {

	var req book_shop.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetBookService().GetBook(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}
