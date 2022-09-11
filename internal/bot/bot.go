package bot

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

const (
	startMessage = "Welcome to PDF Converter\n\nUpload a document and the bot will return the same file to PDF format"
	mpMessage    = "This document cannot be converted to PDF format"
	photoMessage = "Upload a photo as a file"
)

func New(cnv Converter) (*tele.Bot, error) {
	timeout, err := time.ParseDuration(os.Getenv("POLLER_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	settings := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: timeout},
	}

	bot, err := tele.NewBot(settings)
	if err != nil {
		return nil, err
	}

	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send(startMessage)
	})

	bot.Handle(tele.OnAudio, func(ctx tele.Context) error {
		return ctx.Reply(mpMessage)
	})

	bot.Handle(tele.OnVideo, func(ctx tele.Context) error {
		return ctx.Reply(mpMessage)
	})

	bot.Handle(tele.OnPhoto, func(ctx tele.Context) error {
		return ctx.Reply(photoMessage)
	})

	bot.Handle(tele.OnDocument, cnv.Convert)

	return bot, nil
}
