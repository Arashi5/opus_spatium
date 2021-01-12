package logger

type Logger interface {
	log()
	logFatal()
	logPanic()
	customLog()
}
