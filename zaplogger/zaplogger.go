package zaplogger

import (
	"go.uber.org/zap"
)

type (
	WrapperZap struct {
		Z *zap.SugaredLogger
	}
)

func NewZap() (*WrapperZap, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	defer log.Sync()
	return &WrapperZap{
		Z: log.Sugar(),
	}, nil
}

func (z *WrapperZap) Debug(args ...interface{}) {
	z.Z.Debug(args)
}

func (z *WrapperZap) Debugf(template string, args ...interface{}) {
	z.Z.Debugf(template, args)
}

func (z *WrapperZap) Info(args ...interface{}) {
	z.Z.Info(args)
}

func (z *WrapperZap) Infof(template string, args ...interface{}) {
	z.Z.Infof(template, args)
}

func (z *WrapperZap) Warn(args ...interface{}) {
	z.Z.Warn(args)
}

func (z *WrapperZap) Warnf(template string, args ...interface{}) {
	z.Z.Warnf(template, args)
}

func (z *WrapperZap) Error(args ...interface{}) {
	z.Z.Error(args)
}

func (z *WrapperZap) Errorf(template string, args ...interface{}) {
	z.Z.Errorf(template, args)
}

func (z *WrapperZap) DPanic(args ...interface{}) {
	z.Z.DPanic(args)
}

func (z *WrapperZap) DPanicf(template string, args ...interface{}) {
	z.Z.DPanicf(template, args)
}

func (z *WrapperZap) Panic(args ...interface{}) {
	z.Z.Panic(args)
}

func (z *WrapperZap) Panicf(template string, args ...interface{}) {
	z.Z.Panicf(template, args)
}

func (z *WrapperZap) Fatal(args ...interface{}) {
	z.Z.Fatal(args)
}

func (z *WrapperZap) Fatalf(template string, args ...interface{}) {
	z.Z.Fatalf(template, args)
}

func (z *WrapperZap) With(args ...interface{}) *WrapperZap {
	z.Z = z.Z.With(args)
	return z
}
