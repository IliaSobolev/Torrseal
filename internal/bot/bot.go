package bot

import (
	"github.com/IliaSobolev/Torrseal/internal/bot/api"
	"github.com/IliaSobolev/Torrseal/internal/middleware"
	"github.com/IliaSobolev/Torrseal/pkg/domain"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Bot struct {
	tb *telebot.Bot
	l  *layout.Layout
}

func NewBot(tb *telebot.Bot, user domain.UserUsecase, l *layout.Layout) *Bot {
	bot := &Bot{
		tb: tb,
		l:  l,
	}

	m := middleware.NewMiddleware(l, user)

	bot.tb.Use(m.Logger)
	bot.tb.Use(m.User)

	// Routing
	bot.tb.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello, World!")
	})

	bot.tb.Handle(telebot.OnDocument, func(ctx telebot.Context) error {
		return api.HandleFile(ctx)
	})

	bot.tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return api.HandleMagnetLink(ctx)
	})

	return bot
}

func (b *Bot) StartAsync() {
	go b.StartBlocking()
}

func (b *Bot) StartBlocking() {
	_ = b.tb.RemoveWebhook(true)
	b.tb.Start()
}
