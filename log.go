package tgbotapi

import (
	"errors"
	stdlog "log"
	"os"
)

func init() { log = stdlog.New(os.Stderr, "", 0) }

// SetLogger Настройка журналирования для пакета.
func SetLogger(logger BotLogger) (err error) {
	const errNil = "интерфейс журналирования не может быть nil"

	if log = logger; logger == nil {
		err = errors.New(errNil)
		return
	}

	return
}
