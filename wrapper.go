package logme

import (
	"github.com/ervitis/logme/v2/config_loaders"
	"github.com/ervitis/logme/v2/zaplogger"
)

type (
	Wrapper struct {
		log   *zaplogger.WrapperZap
	}
)

func LoggerWrapper(configLoaderModel *config_loaders.LoaderModel) *Wrapper {
	l, _ := zaplogger.NewZap(configLoaderModel)
	return &Wrapper{log: l}
}

func (z *Wrapper) Debug(msg string) {
	z.log.L().Debug(msg)
}

func (z *Wrapper) Info(msg string) {
	z.log.L().Info(msg)
}

func (z *Wrapper) Warn(msg string) {
	z.log.L().Warn(msg)
}

func (z *Wrapper) Error(msg string) {
	z.log.L().Error(msg)
}

func (z *Wrapper) Panic(msg string) {
	z.log.L().Panic(msg)
}
