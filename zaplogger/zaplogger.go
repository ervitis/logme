package zaplogger

import (
	"github.com/ervitis/logme/v2/common"
	"github.com/ervitis/logme/v2/config_loaders"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	WrapperZap struct {
		Z *zap.Logger
	}
)

func NewZap(configModel *config_loaders.LoaderModel) (*WrapperZap, error) {
	level := parseLevel(configModel.Level)

	p := zap.New(zapcore.NewCore(
		parseEncoder(configModel.Encoding),
		zapcore.Lock(common.ParseOutput(configModel.Output)),
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= zapcore.DebugLevel
		}),
	))
	log := p.WithOptions(
		zap.WrapCore(filterLevel(level)),
		zap.Fields(parseFields(configModel.InitialFields)...),
	)
	defer func() {
		_ = log.Sync()
	}()
	return &WrapperZap{
		Z: log,
	}, nil
}

func (z *WrapperZap) L() *zap.Logger {
	return z.Z
}
