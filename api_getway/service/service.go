package service

import (
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	GetUserSevice() book_shop.UserServiceClient
}

type serviceManager struct {
	userService book_shop.UserServiceClient
}

func (s *serviceManager) GetUserSevice() book_shop.UserServiceClient {

	return s.userService
}

func Service() ServiceManagerI {

	userService, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	serviseManager := &serviceManager{userService: book_shop.NewUserServiceClient(userService)}

	return serviseManager
}
