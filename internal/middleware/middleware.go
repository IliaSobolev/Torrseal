package middleware

import "gopkg.in/telebot.v3/layout"

type Middleware struct {
	layout *layout.Layout
}

func NewMiddleware(layout *layout.Layout) *Middleware {
	return &Middleware{
		layout: layout,
	}
}
