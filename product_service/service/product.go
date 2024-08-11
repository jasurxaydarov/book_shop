package service

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type ProductService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedProductServiceServer
}

func NewOrderedItemService(storage storage.StorageRepoI) *ProductService {

	return &ProductService{storage: storage}
}


func (o *ProductService) CreateOrdered_Item(ctx context.Context, req *book_shop.OrderItemCreateReq) (*book_shop.OrderItem, error){
	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().CreateOrderedItem(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}

func (o *ProductService) DeleteOrdered_Item(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error){

	return nil,nil
}
func (o *ProductService) GetOrdered_Item(ctx context.Context,req *book_shop.GetByIdReq) (*book_shop.OrderItem, error){
	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemById(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}
func (o *ProductService) GetOrdered_Items(context.Context, *book_shop.GetListReq) (*book_shop.OrderItemGetListResp, error){

	return nil,nil
}
func (o *ProductService)UpdateOrdered_Item(context.Context, *book_shop.OrderItemCreateReq) (*book_shop.OrderItem, error){

	return nil,nil
}

func (o *ProductService)GetOrdered_ItemByOrderId(ctx context.Context,req *book_shop.GetByIdReq) (*book_shop.OrderItemGetListResp, error){

	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderedItemRepo().GetOrderedItemByOrdreId(ctx, req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}



func (c *ProductService)CreateCategory(ctx context.Context,req *book_shop.CategoryCreateReq) (*book_shop.Category, error){

	log:=logger.NewLogger("",logger.LevelDebug)
	log.Debug("errrrrrrrrr")
	resp, err := c.storage.GetCategoryRepo().CreateCategory(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) GetCategory(ctx context.Context,req *book_shop.GetByIdReq) (*book_shop.Category, error){
	resp, err := c.storage.GetCategoryRepo().GetCategoryById(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *ProductService) GetCategories(context.Context, *book_shop.GetListReq) (*book_shop.CategoryGetListResp, error){
	
	return nil,nil
}

func (c *ProductService) UpdateCategory(context.Context, *book_shop.CategoryUpdateReq) (*book_shop.Category, error){
	
	return nil,nil
}
func (c *ProductService) DeleteCategory(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error){

	return nil,nil
}




func (b *ProductService) CreateBook(ctx context.Context, req *book_shop.BookCreateReq) (*book_shop.Book, error) {

	resp, err := b.storage.GetBookRepo().CreateBook(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (b *ProductService) GetBook(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Book, error) {

	resp, err := b.storage.GetBookRepo().GetBookById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (b *ProductService) DeleteBook(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error) {

	return nil, nil
}

func (b *ProductService) GetBooks(context.Context, *book_shop.GetListReq) (*book_shop.BookGetListResp, error) {

	return nil, nil
}
func (b *ProductService) UpdateBook(context.Context, *book_shop.BookUpdateReq) (*book_shop.Book, error) {

	return nil, nil
}




func (a *ProductService) CreateAuth(ctx context.Context, req *book_shop.AuthorUpdateReq) (*book_shop.Author, error) {

	resp, err := a.storage.GetAuthRepo().CreateAuth(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (a *ProductService) DeleteAuth(ctx context.Context, req *book_shop.DeleteReq) (*book_shop.Empty, error) {

	return nil, nil
}
func (a *ProductService) GetAuth(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Author, error) {

	resp, err := a.storage.GetAuthRepo().GetAuthById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil

}
func (a *ProductService) GetAuths(context.Context, *book_shop.GetListReq) (*book_shop.AuthorGetListResp, error) {

	return nil, nil
}
func (a *ProductService) UpdateAuth(context.Context, *book_shop.AuthorUpdateReq) (*book_shop.User, error) {

	return nil, nil
}


func (o *ProductService) CreateOrder(ctx context.Context, req *book_shop.OrderCreateReq) (*book_shop.Order, error) {

	log:=logger.NewLogger("",logger.LevelDebug)

	log.Debug("errrrrrrrrrrrrr")
	
	resp, err := o.storage.GetOrderRepo().CreateOrder(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) GetOrder(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Order, error) {

	resp, err := o.storage.GetOrderRepo().GetOrderById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (o *ProductService) DeleteOrder(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error) {

	return nil, nil
}

func (o *ProductService) Getorders(context.Context, *book_shop.GetListReq) (*book_shop.OrderGetListResp, error) {

	return nil, nil
}
func (o *ProductService) Updateorder(context.Context, *book_shop.OrderUpdateReq) (*book_shop.Order, error) {

	return nil, nil
}
