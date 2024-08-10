package services

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type UserService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedUserServiceServer
}

func NewUserService(storage storage.StorageRepoI) *UserService {

	return &UserService{storage: storage}
}

func (u *UserService) CreateUser(ctx context.Context, req *book_shop.UserCreateReq) (*book_shop.User, error) {
	log := logger.NewLogger("", logger.LevelDebug)

	resp, err := u.storage.GetUserRepo().CreateUser(context.Background(), req)
	log.Debug("err  CreateUser(ctx ")

	if err != nil {
		log.Debug("err  CreateUser(ctx context.Context,")
		return nil, err
	}
	return resp, nil
}

func (u *UserService) GetUser(context.Context, *book_shop.GetByIdReq) (*book_shop.User, error) {

	return nil, nil
}
func (u *UserService) GetUsers(context.Context, *book_shop.GetListReq) (*book_shop.UserGetListResp, error) {

	return nil, nil
}
func (u *UserService) UpdateUser(context.Context, *book_shop.UserUpdateReq) (*book_shop.User, error) {

	return nil, nil
}

func (u *UserService) DeleteUser(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error) {

	return nil, nil
}
