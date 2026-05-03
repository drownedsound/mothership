package infra

type Logger interface {
	Log(msg string)
}

type LoggerAdapter func(msg string)

func (l LoggerAdapter) Log(msg string) {
	l(msg)
}
