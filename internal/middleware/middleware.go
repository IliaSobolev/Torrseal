package middleware

import (
	"github.com/IliaSobolev/Torrseal/pkg/domain"
	"gopkg.in/telebot.v3/layout"
)

type Middleware struct {
	user   domain.UserUsecase
	layout *layout.Layout
}

func NewMiddleware(layout *layout.Layout, user domain.UserUsecase) *Middleware {
	return &Middleware{
		user:   user,
		layout: layout,
	}
}
