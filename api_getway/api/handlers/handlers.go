package handlers

import (
	"github.com/jasurxaydarov/book_shop/service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Handler struct {
	service service.ServiceManagerI
	log logger.LoggerI
}

func NewHandlers(service service.ServiceManagerI,log logger.LoggerI) Handler {

	return Handler{service: service,}
}
