package ErrChan

import "github.com/skwizi4/lib/logs"

type ErrChan interface {
	InitErrChan(bufferSize int, logger logs.GoLogger) *ErrorChannel
	Start()
	HandleError(err error)
	Close()
}
