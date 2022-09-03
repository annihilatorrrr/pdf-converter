package bot

import (
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

const startMessage = "Welcome to PDF Converter!\n\nSubmit the document and the bot will return the same PDF file"

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

	bot.Use(middleware.Logger())

	bot.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send(startMessage)
	})

	// TODO
	bot.Handle(tele.OnText, cnv.Convert)

	return bot, nil
}
