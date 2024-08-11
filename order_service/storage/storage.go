package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

type StorageRepoI interface{
	GetOrderRepo()postgres.OrderRepoI
}

type storageRepo struct{
	ordeRepo postgres.OrderRepoI
}

func NewOrderRepo(db *pgx.Conn,log logger.LoggerI)StorageRepoI{

	return &storageRepo{ordeRepo: postgres.NewOrderRepo(db,log)}
}


func (s *storageRepo)GetOrderRepo()postgres.OrderRepoI{

	return s.ordeRepo
}
