package service

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

type Logger struct{}

func (lgr *Logger) Error(ctx tele.Context, msg string) {
	log.Printf("| error | [%s]: %s", ctx.Sender().Username, msg)
}

func (lgr *Logger) Info(ctx tele.Context, msg string) {
	log.Printf("| info | [%s]: %s", ctx.Sender().Username, msg)
}
