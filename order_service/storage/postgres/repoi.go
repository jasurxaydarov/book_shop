package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)

type OrderRepoI interface {
	CreateOrder(ctx context.Context, req *book_shop.OrderCreateReq) (*book_shop.Order, error)
	GetOrderById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Order, error)
}
 