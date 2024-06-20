package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Log struct {
	logger *logrus.Logger
}

func InitLogger() GoLogger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	return Log{
		logger: logger,
	}
}

func (l Log) Info(info string) {
	l.logger.Println("INFO:", info)
}
func (l Log) Debug(s string) {
	l.logger.Println("DEBUG:", s)
}

func (l Log) Error(s string) {
	l.logger.Println("ERROR:", s)
}

func (l Log) Fatal(s string) {
	l.logger.Println("FATAL:", s)
}

func (l Log) InfoFrmt(s string, i ...interface{}) {
	str := fmt.Sprintf(s, i...)
	l.Info(str)
}

func (l Log) DebugFrmt(s string, i ...interface{}) {
	str := fmt.Sprintf(s, i...)
	l.Debug(str)
}

func (l Log) ErrorFrmt(s string, i ...interface{}) {
	str := fmt.Sprintf(s, i...)
	l.Error(str)
}

func (l Log) FatalFrmt(s string, i ...interface{}) {
	str := fmt.Sprintf(s, i...)
	l.Fatal(str)
}
