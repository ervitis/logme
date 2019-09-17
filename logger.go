package logme

import (
	loader "github.com/ervitis/logme/config_loaders"
	"github.com/sirupsen/logrus"
)

type Logme struct {
	log  *logrus.Entry
	cfg  loader.ConfigLoader
}

const (
	traceField     = "traceId"
	componentField = "component"
	serviceField   = "service"
)

type Loggerme interface {
	Debug(message, traceId string)
	Info(message, traceId string)
	Warn(message, traceId string)
	Error(message, traceId string)
	AddHook(hook logrus.Hook)
	addFields(fields map[string]interface{}) logrus.Fields
}

func NewLogme(cfg loader.ConfigLoader) *Logme {
	l := logrus.New()
	l.SetOutput(cfg.GetOutput())
	l.SetLevel(cfg.GetLogLevel())
	l.SetFormatter(cfg.GetOutputFormatter())

	return &Logme{
		log: logrus.NewEntry(l),
		cfg: cfg,
	}
}

func (l *Logme) AddHook(hook logrus.Hook) {
	l.log.Logger.Hooks.Add(hook)
}

func (l *Logme) Debug(message, traceId string) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceId})).Debug(message)
}

func (l *Logme) Info(message, traceId string) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceId})).Info(message)
}

func (l *Logme) Warn(message, traceId string) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceId})).Warn(message)
}

func (l *Logme) Error(message, traceId string) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceId})).Error(message)
}

func (l *Logme) addFields(fields map[string]interface{}) logrus.Fields {
	return logrus.Fields{
		traceField:     fields[traceField],
		componentField: l.cfg.GetFixedFields()[componentField],
		serviceField:   l.cfg.GetFixedFields()[serviceField],
	}
}
