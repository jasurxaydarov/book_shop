package service

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
)

type CategoryService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedCategoryServiceServer
}

func NewCategoryService(storage storage.StorageRepoI) *CategoryService {

	return &CategoryService{storage: storage}
}

func (c *CategoryService)CreateCategory(ctx context.Context,req *book_shop.CategoryCreateReq) (*book_shop.Category, error){

	log:=logger.NewLogger("",logger.LevelDebug)
	log.Debug("errrrrrrrrr")
	resp, err := c.storage.GetCategoryRepo().CreateCategory(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *CategoryService) GetCategory(ctx context.Context,req *book_shop.GetByIdReq) (*book_shop.Category, error){
	resp, err := c.storage.GetCategoryRepo().GetCategoryById(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func (c *CategoryService) GetCategories(context.Context, *book_shop.GetListReq) (*book_shop.CategoryGetListResp, error){
	
	return nil,nil
}

func (c *CategoryService) UpdateCategory(context.Context, *book_shop.CategoryUpdateReq) (*book_shop.Category, error){
	
	return nil,nil
}
func (c *CategoryService) DeleteCategory(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error){

	return nil,nil
}