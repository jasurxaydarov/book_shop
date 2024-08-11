package service

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type OrderService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedOrderServiceServer
}

func NewOrderService(storage storage.StorageRepoI) *OrderService {

	return &OrderService{storage: storage}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *book_shop.OrderCreateReq) (*book_shop.Order, error) {

	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderRepo().CreateOrder(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *OrderService) GetOrder(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Order, error) {

	resp, err := o.storage.GetOrderRepo().GetOrderById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *OrderService) DeleteOrder(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error) {

	return nil, nil
}

func (o *OrderService) Getorders(context.Context, *book_shop.GetListReq) (*book_shop.OrderGetListResp, error) {

	return nil, nil
}
func (o *OrderService) Updateorder(context.Context, *book_shop.OrderUpdateReq) (*book_shop.Order, error) {

	return nil, nil
}
