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
 