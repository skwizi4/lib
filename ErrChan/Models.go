package ErrChan

import (
	"main.go/pkg/logs"
)

// Channel - структура для обработки ошибок

type ErrorChannel struct {
	errChan chan error
	logger  logs.GoLogger
}
