package main

import (
	"fmt"

	"github.com/jasurxaydarov/book_shop/api"
	"github.com/jasurxaydarov/book_shop/service"
)


func main(){

	service:=service.Service()

	fmt.Println(service)

	engine:=api.Api(api.Options{Service: service})

	engine.Run(":8080")
}