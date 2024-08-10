package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *Handler) CreateUser(ctx *gin.Context) {

	log:=logger.NewLogger("",logger.LevelDebug)
	var req book_shop.UserCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetUserSevice().CreateUser(context.Background(),&req)

	if err != nil {

		log.Debug("errrrr ")

		log.Debug(err.Error())
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}
