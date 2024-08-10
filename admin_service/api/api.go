package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop/api/handlers"
	"github.com/jasurxaydarov/book_shop/service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Options struct {
	Service service.ServiceManagerI
	log logger.LoggerI
}


func Api(o Options)*gin.Engine{

	h:=handlers.NewHandlers(o.Service,o.log)

	engine:=gin.Default()

	api:=engine.Group("/api")

	api.POST("/create_user",h.CreateUser)

	return engine

}