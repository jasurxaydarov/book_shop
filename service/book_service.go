package service

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type BookService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedBookServiceServer
}

func NewBookService(storage storage.StorageRepoI) *BookService {

	return &BookService{storage: storage}
}
func (b* BookService) CreateBook(ctx context.Context,req *book_shop.BookCreateReq) (*book_shop.Book, error){

	log:=logger.NewLogger("",logger.LevelDebug)
	log.Debug("errrrrrrrrr")
	resp, err := b.storage.GetBookRepo().CreateBook(ctx,req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}

func (b* BookService) GetBook(ctx context.Context,req *book_shop.GetByIdReq) (*book_shop.Book, error){

	log:=logger.NewLogger("",logger.LevelDebug)
	log.Debug("errrrrrrrrr")
	resp, err := b.storage.GetBookRepo().GetBookById(ctx,req)
	
	if err != nil {
	
		fmt.Println(err)
		return nil, err
	}
	
	return resp, nil
}

func (b* BookService) DeleteBook(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error){

	return nil,nil
}

func (b* BookService) GetBooks(context.Context, *book_shop.GetListReq) (*book_shop.BookGetListResp, error){

	return nil,nil
}
func (b* BookService) UpdateBook(context.Context, *book_shop.BookUpdateReq) (*book_shop.Book, error){

	return nil,nil
}


