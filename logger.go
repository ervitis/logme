package logme

import (
	loader "github.com/ervitis/logme/config_loaders"
	"github.com/sirupsen/logrus"
	"sync"
)

type (
	Logme struct {
		log         *logrus.Entry
		cfg         loader.ConfigLoader
		cacheFields logrus.Fields
		mutex       sync.Mutex
		metadata    *Metadata
	}

	Metadata struct {
		TraceID string
	}

	Hook logrus.Hook
)

const (
	traceField = "traceId"
)

type Loggerme interface {
	Debug(message string, metadata ...Metadata)
	Info(message string, metadata ...Metadata)
	Warn(message string, metadata ...Metadata)
	Error(message string, metadata ...Metadata) *logrus.Entry
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
		log: logrus.NewEntry(l),
		cfg: cfg,
	}
}

func (l *Logme) AddHook(hook logrus.Hook) {
	l.log.Logger.Hooks.Add(hook)
}

func (l *Logme) getMetadata(metadata ...Metadata) {
	l.mutex.Lock()
	l.metadata = &Metadata{}
	if len(metadata) > 0 {
		l.metadata = &metadata[0]
	}
	l.mutex.Unlock()
}

func (l *Logme) Debug(message string, metadata ...Metadata) {
	l.getMetadata(metadata...)
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: l.metadata.TraceID})).Debug(message)
}

func (l *Logme) Info(message string, metadata ...Metadata) {
	l.getMetadata(metadata...)
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: l.metadata.TraceID})).Info(message)
}

func (l *Logme) Warn(message string, metadata ...Metadata) {
	l.getMetadata(metadata...)
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: l.metadata.TraceID})).Warn(message)
}

func (l *Logme) Error(message string, metadata ...Metadata) *logrus.Entry {
	l.getMetadata(metadata...)
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: l.metadata.TraceID})).Error(message)
	return l.log
}

func (l *Logme) addFields(fields map[string]interface{}) logrus.Fields {
	if l.cacheFields == nil {
		l.cacheFields = make(logrus.Fields, len(fields)+len(l.cfg.GetFixedFields()))
	}

	for k, v := range fields {
		if _, ok := l.cacheFields[k]; !ok && v != "" {
			l.cacheFields[k] = v
		}
	}
	for k, v := range l.cfg.GetFixedFields() {
		if _, ok := l.cacheFields[k]; !ok {
			l.cacheFields[k] = v
		}
	}
	return l.cacheFields
}

func (l *Logme) L() *logrus.Entry {
	return l.log
}
