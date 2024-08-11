package service

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
)

type AuthService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedAuthServiceServer
}

func NewAuthService(storage storage.StorageRepoI) *AuthService {

	return &AuthService{storage: storage}
}

func (a *AuthService) CreateAuth(ctx context.Context, req *book_shop.AuthorUpdateReq) (*book_shop.Author, error) {

	resp, err := a.storage.GetAuthRepo().CreateAuth(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
func (a *AuthService) DeleteAuth(ctx context.Context, req *book_shop.DeleteReq) (*book_shop.Empty, error) {

	return nil, nil
}
func (a *AuthService) GetAuth(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Author, error) {

	resp, err := a.storage.GetAuthRepo().GetAuthById(ctx, req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil

}
func (a *AuthService) GetAuths(context.Context, *book_shop.GetListReq) (*book_shop.AuthorGetListResp, error) {

	return nil, nil
}
func (a *AuthService) UpdateAuth(context.Context, *book_shop.AuthorUpdateReq) (*book_shop.User, error) {

	return nil, nil
}
