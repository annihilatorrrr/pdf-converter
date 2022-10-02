package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "go.uber.org/automaxprocs"

	"github.com/ew1l/pdf-converter/internal/bot"
	"github.com/ew1l/pdf-converter/internal/service"
)

func main() {
	converter := service.New(&service.Logger{})
	bot, err := bot.New(converter)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		bot.Start()
	}()
	defer bot.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
}
