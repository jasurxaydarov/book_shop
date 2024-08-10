package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)

type CategoryRepoI interface {
	CreateCategory(ctx context.Context, req *book_shop.CategoryCreateReq) (*book_shop.Category, error)
	GetCategoryById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Category, error)
}
 