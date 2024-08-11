package postgres

import (
	"context"

)
type OrderRepoI interface {
	CreateOrder(ctx context.Context, req *book_shop.) (*book_shop.Order, error)
	GetOrderById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Order, error)
}
 