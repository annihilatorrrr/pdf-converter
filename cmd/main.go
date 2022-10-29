package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "go.uber.org/automaxprocs"

	"github.com/ew1l/pdf-converter/internal/bot"
	"github.com/ew1l/pdf-converter/internal/service"
	"github.com/ew1l/pdf-converter/pkg/logger"
)

func main() {
	converter := service.New()
	bot, err := bot.New(converter)
	if err != nil {
		logger.Fatal(err.Error())
	}

	go func() {
		bot.Start()
	}()
	defer bot.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)

	<-stop
}
