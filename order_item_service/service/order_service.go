package service

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type OrderedItemService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedOrderedItemServiceServer
}

func NewOrderedItemService(storage storage.StorageRepoI) *OrderedItemService {

	return &OrderedItemService{storage: storage}
}


func (o *OrderedItemService) CreateOrdered_Item(ctx context.Context, req *book_shop.OrderItemCreateReq) (*book_shop.OrderItem, error){
	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().CreateOrderedItem(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}

func (o *OrderedItemService) DeleteOrdered_Item(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error){

	return nil,nil
}
func (o *OrderedItemService) GetOrdered_Item(ctx context.Context,req *book_shop.GetByIdReq) (*book_shop.OrderItem, error){
	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemById(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}
func (o *OrderedItemService) GetOrdered_Items(context.Context, *book_shop.GetListReq) (*book_shop.OrderItemGetListResp, error){

	return nil,nil
}
func (o *OrderedItemService)UpdateOrdered_Item(context.Context, *book_shop.OrderItemCreateReq) (*book_shop.OrderItem, error){

	return nil,nil
}

func (o *OrderedItemService)GetOrdered_ItemByOrderId(ctx context.Context,req *book_shop.GetByIdReq) (*book_shop.OrderItemGetListResp, error){

	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemByOrdreId(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}