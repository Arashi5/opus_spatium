package logger

type Logger interface {
	Log()
	LogFatal()
	LogPanic()
	CustomLog()
}
