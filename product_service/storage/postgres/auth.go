package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/saidamir98/udevs_pkg/logger"
)

type AuthRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewAuthRepo(db *pgx.Conn, log logger.LoggerI) AuthRepoI {

	return &AuthRepo{db: db, log: log}
}

func (u *AuthRepo) CreateAuth(ctx context.Context, req *book_shop.AuthorUpdateReq) (*book_shop.Author, error) {
	id := uuid.New()
	query := `
		INSERT INTO
			authors (
				author_id,
				name,
				bio 
			)VALUES(
				$1,$2,$3
			)
			`

	_, err := u.db.Exec(
		ctx,
		query,
		id,
		req.AuthorName,
		req.Bio,
	)
	if err != nil {

		u.log.Error("err on db CreateAuth", logger.Error(err))
		return nil, err
	}

	resp, err := u.GetAuthById(context.Background(), &book_shop.GetByIdReq{Id: id.String()})

	if err != nil {

		u.log.Error("err on db GetAuthById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (u *AuthRepo) GetAuthById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.Author, error) {

	var resp book_shop.Author
	qury := `
		SELECT 
			*
		FROM 
			authors 
		WHERE
			author_id = $1
	`

	err := u.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.AuthorId,
		&resp.AuthorName,
		&resp.Bio,
	)

	if err != nil {

		u.log.Error("err on db GetAuthById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}
