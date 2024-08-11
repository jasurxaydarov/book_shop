package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

type StorageRepoI interface{
	GetOrderedItemRepo()postgres.OrderedItemRepoI
}

type storageRepo struct{
	orderedItemRepo postgres.OrderedItemRepoI
}

func NewOrderedItemRepo(db *pgx.Conn,log logger.LoggerI)StorageRepoI{

	return &storageRepo{orderedItemRepo: postgres.NewOrderedItemRepo(db,log)}
}


func (s *storageRepo)GetOrderedItemRepo()postgres.OrderedItemRepoI{

	return s.orderedItemRepo
}
