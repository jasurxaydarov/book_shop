package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)

type AuthRepoI interface {
	CreateAuth(ctx context.Context, req *book_shop.AuthorUpdateReq) (*book_shop.Author, error)
	GetAuthById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Author, error)
}
 