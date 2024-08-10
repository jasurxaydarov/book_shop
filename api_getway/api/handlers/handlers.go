package handlers

import (
	"github.com/jasurxaydarov/book_shop/service"
)

type Handler struct {
	service service.ServiceManagerI
}

func NewHandlers(service service.ServiceManagerI) Handler {

	return Handler{service: service}
}
