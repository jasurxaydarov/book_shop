package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/pkg/db"
	"github.com/jasurxaydarov/book_shop/service"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
)

func main() {

	log := logger.NewLogger("", logger.LevelDebug)

	pgxConn, err := db.ConnDB(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pgxConn)

	storage:=storage.NewBookRepo(pgxConn,log)

	service:=service.NewBookService(storage)

	listen,err:=net.Listen("tcp","localhost:8003")

	if err != nil {
		fmt.Println(err)
		return
	}

	server:=grpc.NewServer()

	book_shop.RegisterBookServiceServer(server,service)

	log.Debug("server serve on :8003")

	if err =server.Serve(listen);err!=nil{
		log.Error(err.Error())
		return 

	}
}
