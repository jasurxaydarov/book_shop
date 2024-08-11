package main

import (
	"fmt"

	"github.com/jasurxaydarov/book_shop/api"
	"github.com/jasurxaydarov/book_shop/service"
	"github.com/saidamir98/udevs_pkg/logger"
)


func main(){

	log:=logger.NewLogger("",logger.LevelDebug)
	service:=service.Service()

	fmt.Println(service)

	engine:=api.Api(api.Options{Service: service,Log:log,})

	engine.Run(":8080")
}