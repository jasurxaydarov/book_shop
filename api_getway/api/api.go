package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop/api/handlers"
	"github.com/jasurxaydarov/book_shop/service"
)

type Options struct {
	Service service.ServiceManagerI
}

func Api(o Options) *gin.Engine {

	h := handlers.NewHandlers(o.Service)

	engine := gin.Default()

	api := engine.Group("/api")

	api.POST("/create_user", h.CreateUser)

	return engine

}
