package logme

import (
	loader "github.com/ervitis/logme/config_loaders"
	"github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

type (
	Logme struct {
		log      *logrus.Entry
		mutex    sync.Mutex
		metadata *Metadata
	}

	Metadata struct {
		TraceID string `logme:"traceId"`
	}

	Hook logrus.Hook
)

type Loggerme interface {
	L(metadata ...Metadata) *logrus.Entry
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

func (l *Logme) getMetadata(metadata ...Metadata) {
	l.mutex.Lock()
	l.metadata = &Metadata{}
	if len(metadata) > 0 {
		l.metadata = &metadata[0]
	}
	l.mutex.Unlock()
}

func (l *Logme) L(metadata ...Metadata) *logrus.Entry {
	l.getMetadata(metadata...)
	return l.log.WithFields(l.addFields())
}

func (l *Logme) addFields() logrus.Fields {
	rv := reflect.ValueOf(l.metadata)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	fields := make(map[string]interface{}, 0)

	for i := 0; i < rv.NumField(); i++ {
		if rv.Field(i).IsZero() {
			continue
		}

		t := rv.Type().Field(i).Tag.Get("logme")
		if t == "" {
			t = rv.Type().Field(i).Name
		}
		fields[t] = rv.Field(i).Interface()
	}

	return fields
}
