package mongo

import (
	"context"
	"errors"
	"github.com/IliaSobolev/Torrseal/pkg/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const GetQuery = "SELECT * FROM users WHERE id = $1"
const CreateQuery = "INSERT INTO users (id, lang) VALUES ($1, $2) RETURNING users.id, users.lang"

type repo struct {
	conn *pgxpool.Pool
}

func NewUserRepo(conn *pgxpool.Pool) domain.UserRepository {
	return &repo{conn: conn}
}

func (u *repo) Get(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	err := u.conn.QueryRow(ctx, GetQuery, id).Scan(&user.ID, &user.Lang, &user.ExpiryDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
	}
	return &user, nil
}

func (u *repo) Create(ctx context.Context, obj *domain.User) (*domain.User, error) {
	var user domain.User
	err := u.conn.QueryRow(ctx, CreateQuery, obj.ID, obj.Lang).Scan(&user.ID, &user.Lang)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
