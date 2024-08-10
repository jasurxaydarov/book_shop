package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)


type UserRepoI interface{

	CreateUser(ctx context.Context, req *book_shop.UserCreateReq)(*book_shop.User,error)
	GetUserById(ctx context.Context, req *book_shop.GetByIdReq)(*book_shop.User,error)

}