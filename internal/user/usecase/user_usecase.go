package usecase

import "github.com/IliaSobolev/Torrseal/pkg/domain"

type uc struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &uc{repo: repo}
}
