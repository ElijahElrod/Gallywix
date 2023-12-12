package logger

// Logger is an interface of log applications
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}
