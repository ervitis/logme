package logme

import (
	loader "github.com/ervitis/logme/config_loaders"
	"github.com/sirupsen/logrus"
)

type (
	Logme struct {
		log      *logrus.Entry
	}

	Hook logrus.Hook
)

type Loggerme interface {
	L() *logrus.Entry
}

func NewLogme(cfg loader.ConfigLoader, hooks ...Hook) Loggerme {
	l := logrus.New()
	l.SetOutput(cfg.GetOutput())
	l.SetLevel(cfg.GetLogLevel())
	l.SetFormatter(cfg.GetOutputFormatter())

	for _, hook := range hooks {
		l.AddHook(hook)
	}

	return &Logme{
		log: logrus.NewEntry(l).WithFields(cfg.GetFixedFields()),
	}
}

func (l *Logme) AddHook(hook logrus.Hook) {
	l.log.Logger.Hooks.Add(hook)
}

func (l *Logme) L() *logrus.Entry {
	return l.log
}
