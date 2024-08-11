package service

import (
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Service() ServiceManagerI {

	userService, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	authService, err := grpc.NewClient("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	categoryService, err := grpc.NewClient("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	bookService, err := grpc.NewClient("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	oederService, err := grpc.NewClient("localhost:8004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	orderItemService, err := grpc.NewClient("localhost:8005", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	serviseManager := &serviceManager{
		userService:     book_shop.NewUserServiceClient(userService),
		authService:     book_shop.NewAuthServiceClient(authService),
		categoryService: book_shop.NewCategoryServiceClient(categoryService),
		bookService:     book_shop.NewBookServiceClient(bookService),
		orderService:    book_shop.NewOrderServiceClient(oederService),
		orderItemService: book_shop.NewOrderedItemServiceClient(orderItemService),
	}

	return serviseManager
}

type ServiceManagerI interface {
	GetUserSevice() book_shop.UserServiceClient
	GetAuthService() book_shop.AuthServiceClient
	GetCategoryService() book_shop.CategoryServiceClient
	GetBookService() book_shop.BookServiceClient
	GetOrderService() book_shop.OrderServiceClient
	GetOrderItemService() book_shop.OrderedItemServiceClient

}

type serviceManager struct {
	userService      book_shop.UserServiceClient
	authService      book_shop.AuthServiceClient
	categoryService  book_shop.CategoryServiceClient
	bookService      book_shop.BookServiceClient
	orderService     book_shop.OrderServiceClient
	orderItemService book_shop.OrderedItemServiceClient
}

func (s *serviceManager) GetUserSevice() book_shop.UserServiceClient {

	return s.userService
}

func (s *serviceManager) GetAuthService() book_shop.AuthServiceClient {

	return s.authService
}

func (s *serviceManager) GetCategoryService() book_shop.CategoryServiceClient {

	return s.categoryService
}

func (s *serviceManager) GetBookService() book_shop.BookServiceClient {

	return s.bookService
}

func (s *serviceManager) GetOrderService() book_shop.OrderServiceClient {

	return s.orderService
}

func (s *serviceManager) GetOrderItemService() book_shop.OrderedItemServiceClient{

	return s.orderItemService
}

