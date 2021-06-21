package logme

import (
	"github.com/ervitis/logme/v2/zaplogger"
)

type (
	Wrapper struct {
		Log *zaplogger.WrapperZap
	}
)

func LoggerWrapper() *Wrapper {
	l, _ := zaplogger.NewZap()
	return &Wrapper{Log: l}
}

func (z *Wrapper) Debug(args ...interface{}) {
	z.Log.Debug(args)
}

func (z *Wrapper) Debugf(template string, args ...interface{}) {
	z.Log.Debugf(template, args)
}

func (z *Wrapper) Info(args ...interface{}) {
	z.Log.Info(args)
}

func (z *Wrapper) Infof(template string, args ...interface{}) {
	z.Log.Infof(template, args)
}

func (z *Wrapper) Warn(args ...interface{}) {
	z.Log.Warn(args)
}

func (z *Wrapper) Warnf(template string, args ...interface{}) {
	z.Log.Warnf(template, args)
}

func (z *Wrapper) Error(args ...interface{}) {
	z.Log.Error(args)
}

func (z *Wrapper) Errorf(template string, args ...interface{}) {
	z.Log.Errorf(template, args)
}

func (z *Wrapper) Panic(args ...interface{}) {
	z.Log.Panic(args)
}

func (z *Wrapper) Panicf(template string, args ...interface{}) {
	z.Log.Panicf(template, args)
}

func (z *Wrapper) With(fields ...interface{}) *Wrapper {
	z.Log = z.Log.With(fields...)
	return z
}
