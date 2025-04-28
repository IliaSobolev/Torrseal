package bot

import (
	"github.com/IliaSobolev/Torrseal/internal/middleware"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/layout"
)

type Bot struct {
	tb *telebot.Bot
	l  *layout.Layout
}

func NewBot(tb *telebot.Bot, l *layout.Layout) *Bot {
	b := &Bot{
		tb: tb,
		l:  l,
	}

	m := middleware.NewMiddleware(l)

	b.tb.Use(m.Logger)

	//Test
	b.tb.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello, World!")
	})

	b.tb.Handle(telebot.OnDocument, func(ctx telebot.Context) error {
		return handleFile(ctx)
	})

	b.tb.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return handleMagnetLink(ctx)
	})

	return b
}

func (b *Bot) StartAsync() {
	go b.StartBlocking()
}

func (b *Bot) StartBlocking() {
	_ = b.tb.RemoveWebhook(true)
	b.tb.Start()
}
