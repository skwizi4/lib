package ErrChan

import "main.go/pkg/logs"

type ErrChan interface {
	InitErrChan(bufferSize int, logger logs.GoLogger) *ErrorChannel
	Start()
	HandleError(err error)
	Close()
}
