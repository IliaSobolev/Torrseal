package usecase

import (
	"context"
	"github.com/IliaSobolev/Torrseal/pkg/domain"
	"github.com/pkg/errors"
)

func (uc *uc) Get(ctx context.Context, id int64) (*domain.User, error) {
	user, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "user.repo.Get")
	}
	return user, nil
}
