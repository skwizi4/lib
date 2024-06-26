package ErrChan

import (
	logger "github.com/skwizi4/lib/logs"
)

func InitErrChan(bufferSize int, logger logger.GoLogger) *ErrorChannel {
	var ChanelErr ErrorChannel
	ChanelErr.logger = logger
	ChanelErr.errChan = make(chan error, bufferSize)
	return &ChanelErr
}

// Start - запуск горутины для обработки ошибок
func (e *ErrorChannel) Start() {
	go func() {
		for err := range e.errChan {
			e.logger.ErrorFrmt("ERROR : %s", err)
		}
	}()
}

// HandleError - отправка ошибки в канал
func (e *ErrorChannel) HandleError(err error) {
	if err != nil {
		e.errChan <- err
	}
}

// Close - закрытие канала ошибок
func (e *ErrorChannel) Close() {
	close(e.errChan)
}
