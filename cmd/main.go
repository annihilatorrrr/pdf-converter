package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ew1l/pdf-converter/internal/bot"
	"github.com/ew1l/pdf-converter/internal/service"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	bot, err := bot.New(&service.Converter{})
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
