package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

type StorageRepoI interface{
	GetCategoryRepo()postgres.CategoryRepoI
}

type storageRepo struct{
	categoryRepo postgres.CategoryRepoI
}

func NewCategoryRepo(db *pgx.Conn,log logger.LoggerI)StorageRepoI{

	return &storageRepo{categoryRepo: postgres.NewCategoryRepo(db,log)}
}


func (s *storageRepo)GetCategoryRepo()postgres.CategoryRepoI{

	return s.categoryRepo
}
