package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

type StorageRepoI interface{
	GetAuthRepo()postgres.AuthRepoI
}

type storageRepo struct{
	authRepo postgres.AuthRepoI
}

func NewAuthRepo(db *pgx.Conn,log logger.LoggerI)StorageRepoI{

	return &storageRepo{authRepo: postgres.NewAuthRepo(db,log)}
}


func (s *storageRepo)GetAuthRepo()postgres.AuthRepoI{

	return s.authRepo
}
