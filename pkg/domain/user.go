package domain

import (
	"context"
	"errors"
	"time"
)

type User struct {
	ID         int64
	Lang       string
	ExpiryDate time.Time
}

var Locales = []string{"en", "ru"}

var ErrUserNotFound = errors.New("user not found")

type UserUsecase interface {
	Get(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, id int64) (*User, error)
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}
