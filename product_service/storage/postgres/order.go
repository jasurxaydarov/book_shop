package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/saidamir98/udevs_pkg/logger"
)

type orderRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewOrderRepo(db *pgx.Conn, log logger.LoggerI) OrderRepoI {

	return &orderRepo{db: db, log: log}
}

func (o *orderRepo) CreateOrder(ctx context.Context, req *book_shop.OrderCreateReq) (*book_shop.Order, error) {
	id := uuid.New()

	query := `
		INSERT INTO
			orders(
				order_id,
				user_id,
				total_amount,
				order_status 
			)VALUES(
				$1,$2,$3,$4
			)
			`

	_, err := o.db.Exec(
		ctx,
		query,
		id,
		req.UserId,
		req.TotalAmount,
		req.OrderStatus,
	)
	if err != nil {

		o.log.Error("err on db CreateOrder", logger.Error(err))
		return nil, err
	}

	resp, err := o.GetOrderById(context.Background(), &book_shop.GetByIdReq{Id: id.String()})

	if err != nil {

		o.log.Error("err on db GetOrderById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (o *orderRepo) GetOrderById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Order, error) {

	var resp book_shop.Order
	qury := `
		SELECT 
			order_id,
			user_id,
			total_amount,
			order_status 
		FROM 
			orders
		WHERE
			order_id = $1
	`

	err := o.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.OrderId,
		&resp.UserId,
		&resp.TotalAmount,
		&resp.OrderStatus,
	
	)

	if err != nil {

		o.log.Error("err on db GetBookById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}