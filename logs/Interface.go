package logs

type GoLogger interface {
	Info(string)
	Debug(string)
	Error(string)
	Fatal(string)
	InfoFrmt(string, ...interface{})
	DebugFrmt(string, ...interface{})
	ErrorFrmt(string, ...interface{})
	FatalFrmt(string, ...interface{})
}
