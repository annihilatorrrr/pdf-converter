package service

import tele "gopkg.in/telebot.v3"

type Converter struct{}

func (conv *Converter) Convert(ctx tele.Context) error {
	return ctx.Send(ctx.Text())
}
