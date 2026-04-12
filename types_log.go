package tgbotapi

// BotLogger Интерфейс методов журнала регистрации.
type BotLogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}
