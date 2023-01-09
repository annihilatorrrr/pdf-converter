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

		ctx.Reply("In process...")

		if err := ctx.Bot().Download(ctx.Message().Document.MediaFile(), input); err != nil {
			cnv.failure(ctx, err, args...)
		}
		defer os.Remove(input)

		if _, err := cnv.unoconv(input); err != nil {
			cnv.failure(ctx, err, args...)
		}
		defer os.Remove(output)

		if err := ctx.Reply(&tele.Document{
			File:     tele.FromDisk(output),
			FileName: output,
		}); err != nil {
			cnv.failure(ctx, err, args...)
		}

		logger.Info("pdfcnv", args...)
	}(ctx)

	return nil
}

func (cnv *Converter) unoconv(input string) ([]byte, error) {
	cnv.Lock()
	defer cnv.Unlock()

	return exec.Command("unoconv", input).Output()
}

func (cnv *Converter) failure(ctx tele.Context, err error, args ...any) {
	logger.Error(err.Error(), args...)
	ctx.Reply("Something went wrong! Please try again!")
}
