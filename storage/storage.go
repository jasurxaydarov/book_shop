package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

type StorageRepoI interface{
	GetBookRepo()postgres.BookRepoI
}

type storageRepo struct{
	bookRepo postgres.BookRepoI
}

func NewBookRepo(db *pgx.Conn,log logger.LoggerI)StorageRepoI{

	return &storageRepo{bookRepo: postgres.NewBookRepo(db,log)}
}


func (s *storageRepo)GetBookRepo()postgres.BookRepoI{

	return s.bookRepo
}
