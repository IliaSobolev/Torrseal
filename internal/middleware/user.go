package middleware

import (
	"context"
	"github.com/IliaSobolev/Torrseal/pkg/domain"
	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
)

func (m *Middleware) User(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		user, err := m.user.Get(context.Background(), c.Sender().ID)
		if errors.Is(err, domain.ErrUserNotFound) {
			user, err = m.user.Create(context.Background(), c.Sender().ID)
			if err != nil {
				return errors.Wrap(err, "user.usecase.Create")
			}
		} else if err != nil {
			return errors.Wrap(err, "user.usecase.Get")
		}

		found := false
		for _, locale := range domain.Locales {
			if locale == user.Lang {
				found = true
				break
			}
		}
		if found {
			m.layout.SetLocale(c, user.Lang)
		} else {
			m.layout.SetLocale(c, "en")
		}
		err = next(c)
		if err != nil {
			return c.Send(m.layout.Text(c, "forbidden"))
		}
		return err
	}
}
