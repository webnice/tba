package tgbotapi

import (
	"context"
	"net/http"
	"sync"
)

var log BotLogger

// BotAPI Объект сущности пакета.
//
//goland:noinspection GoVetStructTag
type BotAPI struct {
	Token        string               `json:"token"`         // Токен бота.
	Debug        bool                 `json:"debug"`         // Флаг режима отладки.
	BufferLength int64                `json:"buffer_length"` // Размер буфера.
	Self         User                 `json:"-"`             // Объект пользователя бота.
	Client       HTTPClient           `json:"-"`             // HTTP клиент.
	endpointApi  string               `json:"-"`             // URI шаблон API телеграм.
	endpointFile string               `json:"-"`             // URI шаблон API доступа к файлам.
	stoppers     []context.CancelFunc `json:"-"`             // Функции прерывания контекста.
	stoppersMux  *sync.RWMutex        `json:"-"`             // Защита от конкурентного доступа для stoppers.
}

// HTTPClient HTTP клиент для выполнения запросов к серверу телеграм.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
