package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/saidamir98/udevs_pkg/logger"
)

type orderedItemRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewOrderedItemRepo(db *pgx.Conn, log logger.LoggerI) OrderedItemRepoI {

	return &orderedItemRepo{db: db, log: log}
}

func (o *orderedItemRepo) CreateOrderedItem(ctx context.Context, req *book_shop.OrderItemCreateReq) (*book_shop.OrderItem, error) {
	id := uuid.New()

	query := `
		INSERT INTO
			order_items(
				order_item_id,
				order_id,
				user_id,
				book_id,
				quantity,
				price 
			)VALUES(
				$1,$2,$3,$4,$5,$6
			)
			`

	_, err := o.db.Exec(
		ctx,
		query,
		id,
		req.OrderId,
		req.UserId,
		req.BookId,
		req.Quantity,
		req.Price,
	)
	if err != nil {

		o.log.Error("err on db CreateOrderItem", logger.Error(err))
		return nil, err
	}

	resp, err := o.GetOrderedItemById(context.Background(), &book_shop.GetByIdReq{Id: id.String()})

	if err != nil {

		o.log.Error("err on db GetOrderItemById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (o *orderedItemRepo) GetOrderedItemById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.OrderItem, error) {

	var resp book_shop.OrderItem
	qury := `
		SELECT 
			order_item_id,
			order_id,
			user_id,
			book_id,
			quantity,
			price 
		FROM 
			order_items
		WHERE
			order_item_id = $1
	`

	err := o.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.OrderItemId,
		&resp.OrderId,
		&resp.UserId,
		&resp.BookId,
		&resp.Quantity,
		&resp.Price,
	)

	if err != nil {

		o.log.Error("err on db GetOrderItemByOrderId scan", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (o *orderedItemRepo)GetOrderedItemByOrdreId(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.OrderItemGetListResp, error){

	var row book_shop.OrderItemGetListResp
	var resp book_shop.OrderItem

	qury := `
		SELECT 
			order_item_id,
			order_id,
			user_id,
			book_id,
			quantity,
			price 
		FROM 
			order_items
		WHERE
			order_id = $1
	`

	rows, err := o.db.Query(ctx, qury, req.Id)

	for rows.Next() {

		rows.Scan(
			&resp.OrderItemId,
			&resp.OrderId,
			&resp.UserId,
			&resp.BookId,
			&resp.Quantity,
			&resp.Price,
		)

		row.OrderItem = append(row.OrderItem, &resp)

		row.Count++
	}

	if err != nil {

		o.log.Error("err on db GetOrderItemByOrderId", logger.Error(err))
		return nil, err
	}

	return &row, nil
}
