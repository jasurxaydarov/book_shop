package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/saidamir98/udevs_pkg/logger"
)

type UserRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewUserRepo(db *pgx.Conn, log logger.LoggerI) UserRepoI {

	return &UserRepo{db: db, log: log}
}

func (u *UserRepo) CreateUser(ctx context.Context, req *book_shop.UserCreateReq) (*book_shop.User, error) {
u.log.Debug("aaaaaaaaaaaaaaaaa")
	id := uuid.New()
	query := `
		INSERT INTO
			users (
				user_id,
				username,
				email, 
				password,
				full_name,
				is_admin
			)VALUES(
				$1,$2,$3,$4,$5,$6
			)
			`

	_, err := u.db.Exec(
		ctx,
		query,
		id,
		req.Username,
		req.Email,
		req.Password,
		req.Fullname,
		req.IsAdmin,
	)
	if err != nil {

		u.log.Debug("errrrrr")
		u.log.Error("err on db CreateUser", logger.Error(err))
		return nil, err
	}

	resp, err := u.GetUserById(context.Background(), &book_shop.GetByIdReq{Id: id.String()})

	if err != nil {

		u.log.Error("err on db GetUserById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.User, error) {

	var resp book_shop.User
	qury := `
		SELECT 
			user_id,
				username,
				email, 
				password,
				full_name,
				is_admin
		FROM 
			users 
		WHERE
			user_id = $1
	`

	err := u.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.UserId,
		&resp.Username,
		&resp.Email,
		&resp.Password,
		&resp.Fullname,
		&resp.IsAdmin,
	
	)

	if err != nil {

		u.log.Error("err on db GetUserById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}
