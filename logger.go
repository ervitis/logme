package logme

import (
	loader "github.com/ervitis/logme/config_loaders"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Logme struct {
	log  *logrus.Entry
	cfg  loader.ConfigLoader
}

const (
	traceField     = "traceid"
	codeField      = "code"
	componentField = "component"
	serviceField   = "service"
)

type Loggerme interface {
	Debug(message, traceid string, code int)
	Info(message, traceid string, code int)
	Warn(message, traceid string, code int)
	Error(message, traceid string)
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

func (l *Logme) Debug(message, traceid string, code int) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceid, codeField: strconv.Itoa(code)})).Debug(message)
}

func (l *Logme) Info(message, traceid string, code int) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceid, codeField: strconv.Itoa(code)})).Info(message)
}

func (l *Logme) Warn(message, traceid string, code int) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceid, codeField: strconv.Itoa(code)})).Warn(message)
}

func (l *Logme) Error(message, traceid string) {
	l.log.WithFields(l.addFields(map[string]interface{}{traceField: traceid, codeField: strconv.Itoa(http.StatusInternalServerError)})).Error(message)
}

func (l *Logme) addFields(fields map[string]interface{}) logrus.Fields {
	return logrus.Fields{
		traceField:     fields[traceField],
		codeField:      fields[codeField],
		componentField: l.cfg.GetFixedFields()[componentField],
		serviceField:   l.cfg.GetFixedFields()[serviceField],
	}
}
