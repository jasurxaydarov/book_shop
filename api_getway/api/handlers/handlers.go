package handlers

import (
	"github.com/jasurxaydarov/book_shop/redis"
	"github.com/jasurxaydarov/book_shop/service"
	"github.com/saidamir98/udevs_pkg/logger"
)

type Handler struct {
	service service.ServiceManagerI
	log     logger.LoggerI
	cache   redis.RedisRepoI
}

func NewHandlers(service service.ServiceManagerI, log logger.LoggerI, cache redis.RedisRepoI) Handler {

	return Handler{service: service ,log: log,cache: cache}
}
