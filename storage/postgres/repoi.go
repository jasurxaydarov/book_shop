package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)

type BookRepoI interface {
	CreateBook(ctx context.Context, req *book_shop.BookCreateReq) (*book_shop.Book, error)
	GetBookById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Book, error)
}
 