package usecase

import (
	"context"
	"github.com/IliaSobolev/Torrseal/pkg/domain"
	"github.com/pkg/errors"
)

func (uc *uc) Create(ctx context.Context, id int64) (*domain.User, error) {
	obj := &domain.User{
		ID:   id,
		Lang: "en",
	}
	user, err := uc.repo.Create(ctx, obj)
	if err != nil {
		return nil, errors.Wrap(err, "user.repo.Create")
	}
	return user, nil
}
