package service

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

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

		if err := ctx.Bot().Download(ctx.Message().Document.MediaFile(), input); err != nil {
			return
		}
		defer os.Remove(input)

		if _, err := cnv.unoconv(input); err != nil {
			return
		}
		defer os.Remove(output)

		if err := ctx.Reply(&tele.Document{
			File:     tele.FromDisk(output),
			FileName: output,
		}); err != nil {
			return
		}
	}(ctx)

	return nil
}

func (cnv *Converter) unoconv(input string) ([]byte, error) {
	cnv.Lock()
	defer cnv.Unlock()

	return exec.Command("unoconv", input).Output()
}
