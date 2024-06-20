package ErrChan

import (
	"github.com/skwizi4/lib/logs"
)

// Channel - структура для обработки ошибок

type ErrorChannel struct {
	errChan chan error
	logger  logs.GoLogger
}
