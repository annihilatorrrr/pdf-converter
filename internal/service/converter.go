package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	tele "gopkg.in/telebot.v3"
)

type Converter struct {
	sync.Mutex
	lgr *Logger
}

func New(lgr *Logger) *Converter {
	return &Converter{
		lgr: lgr,
	}
}

func (cnv *Converter) Convert(ctx tele.Context) error {
	go func(ctx tele.Context) {
		input := ctx.Message().Document.FileName
		output := strings.TrimSuffix(input, filepath.Ext(input)) + ".pdf"

		if err := ctx.Bot().Download(ctx.Message().Document.MediaFile(), input); err != nil {
			cnv.lgr.Error(ctx, err.Error())
			return
		}
		defer os.Remove(input)

		if msg, err := cnv.unoconv(input); err != nil {
			cnv.lgr.Error(ctx, string(msg))
			return
		}
		defer os.Remove(output)

		if err := ctx.Reply(&tele.Document{
			File:     tele.FromDisk(output),
			FileName: output,
		}); err != nil {
			cnv.lgr.Error(ctx, err.Error())
			return
		}

		cnv.lgr.Info(ctx, fmt.Sprintf("%s -> %s", input, output))
	}(ctx)

	return nil
}

func (cnv *Converter) unoconv(input string) ([]byte, error) {
	cnv.Lock()
	defer cnv.Unlock()

	return exec.Command("unoconv", input).Output()
}
