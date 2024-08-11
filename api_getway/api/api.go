package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop/api/handlers"
	"github.com/jasurxaydarov/book_shop/service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Options struct {
	Service service.ServiceManagerI
	Log     logger.LoggerI
}

func Api(o Options) *gin.Engine {

	h := handlers.NewHandlers(o.Service,o.Log)

	engine := gin.Default()

	api := engine.Group("/api")
	us := api.Group("/us")
	{
		us.POST("/user", h.CreateUser)
		us.GET("/user/:id", h.GetUserById)

		//order
		us.POST("/order", h.CreateOrder)
		us.GET("/order/:id", h.GetOrderById)

		//orderItem
		us.POST("/order_item", h.CreateOrderItem)
		us.GET("/order_item/:id", h.GetOrderItemById)
		us.GET("/order_item_id/:id", h.GetOrderItemById)


	}

	adm := api.Group("/adm")

	{
		// author
		adm.POST("/auth", h.CreateAuth)
		adm.GET("/auth/:id", h.GetAuthById)

		//category
		adm.POST("/category", h.CreateCategory)
		adm.GET("/category/:id", h.GetCategoryById)

		//book
		adm.POST("/book", h.CreateBook)
		adm.GET("/book/:id", h.GetBookById)

	}

	return engine

}
