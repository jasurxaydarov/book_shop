package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)

type OrderedItemRepoI interface {
	CreateOrderedItem(ctx context.Context, req *book_shop.OrderItemCreateReq) (*book_shop.OrderItem, error)
	GetOrderedItemById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.OrderItem, error)
	GetOrderedItemByOrdreId(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.OrderItemGetListResp, error)

}
 
type OrderRepoI interface {
	CreateOrder(ctx context.Context, req *book_shop.OrderCreateReq) (*book_shop.Order, error)
	GetOrderById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Order, error)
}
 
type AuthRepoI interface {
	CreateAuth(ctx context.Context, req *book_shop.AuthorUpdateReq) (*book_shop.Author, error)
	GetAuthById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Author, error)
}
 
type BookRepoI interface {
	CreateBook(ctx context.Context, req *book_shop.BookCreateReq) (*book_shop.Book, error)
	GetBookById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Book, error)
}
 
type CategoryRepoI interface {
	CreateCategory(ctx context.Context, req *book_shop.CategoryCreateReq) (*book_shop.Category, error)
	GetCategoryById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Category, error)
}
 