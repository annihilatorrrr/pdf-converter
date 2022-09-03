package bot

import tele "gopkg.in/telebot.v3"

type Converter interface {
	Convert(c tele.Context) error
}
