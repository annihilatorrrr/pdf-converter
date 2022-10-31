package service

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ew1l/pdf-converter/pkg/logger"
	tele "gopkg.in/telebot.v3"
)

type Converter struct {
	sync.Mutex
}

func New() *Converter {
	return &Converter{}
}

func (cnv *Converter) Convert(ctx tele.Context) error {
	go func(ctx tele.Context) {
		input := ctx.Message().Document.FileName
		output := strings.TrimSuffix(input, filepath.Ext(input)) + ".pdf"
		args := []any{"username", ctx.Sender().Username, "input", input, "output", output}

		if err := ctx.Reply("In process..."); err != nil {
			logger.Error(err.Error(), args...)
			return
		}

		if err := ctx.Bot().Download(ctx.Message().Document.MediaFile(), input); err != nil {
			logger.Error(err.Error(), args...)
			if err := cnv.failure(ctx); err != nil {
				logger.Error(err.Error(), args...)
			}
			return
		}
		defer os.Remove(input)

		if _, err := cnv.unoconv(input); err != nil {
			logger.Error(err.Error(), args...)
			if err := cnv.failure(ctx); err != nil {
				logger.Error(err.Error(), args...)
			}
			return
		}
		defer os.Remove(output)

		if err := ctx.Reply(&tele.Document{
			File:     tele.FromDisk(output),
			FileName: output,
		}); err != nil {
			logger.Error(err.Error(), args...)
			if err := cnv.failure(ctx); err != nil {
				logger.Error(err.Error(), args...)
			}
			return
		}

		logger.Info("Сompleted", args...)
	}(ctx)

	return nil
}

func (cnv *Converter) unoconv(input string) ([]byte, error) {
	cnv.Lock()
	defer cnv.Unlock()

	return exec.Command("unoconv", input).Output()
}

func (cnv *Converter) failure(ctx tele.Context) error {
	if err := ctx.Reply("Something went wrong! Please try again"); err != nil {
		return err
	}

	return nil
}
