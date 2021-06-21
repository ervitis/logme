package logme

type (
	Loggerme struct {
		logWrap *Wrapper
	}
)

func NewLogme() *Loggerme {
	return &Loggerme{logWrap: LoggerWrapper()}
}

func (l *Loggerme) L() *Wrapper {
	return l.logWrap
}
