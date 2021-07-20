package zaplogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

type (
	levelFilterCore struct {
		zapcore.Core
		level zapcore.Level
	}
)

func newLevelFilterCore(core zapcore.Core, level zapcore.Level) zapcore.Core {
	return &levelFilterCore{core, level}
}

func (c *levelFilterCore) Enabled(lvl zapcore.Level) bool {
	return lvl >= c.level
}

func (c *levelFilterCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if !c.Enabled(ent.Level) {
		return ce
	}

	return c.Core.Check(ent, ce)
}

func filterLevel(level zapcore.Level) func(core zapcore.Core) zapcore.Core {
	return func(core zapcore.Core) zapcore.Core {
		return newLevelFilterCore(core, level)
	}
}

func parseLevel(level string) zapcore.Level {
	level = strings.ToUpper(level)
	levels := map[string]zapcore.Level {
		"DEBUG": zapcore.DebugLevel,
		"INFO": zapcore.InfoLevel,
		"WARN": zapcore.WarnLevel,
		"ERROR": zapcore.ErrorLevel,
	}
	if lvl, exists := levels[level]; !exists {
		return zapcore.InfoLevel
	} else {
		return lvl
	}
}

func parseEncoder(encoder string) zapcore.Encoder {
	encoder = strings.ToUpper(encoder)
	encoders := map[string]zapcore.Encoder {
		"JSON": zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		"TEXT": zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
	}
	if codec, exists := encoders[encoder]; !exists {
		return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	} else {
		return codec
	}
}

func parseOutput(output string) *os.File {
	output = strings.ToUpper(output)
	outputs := map[string]*os.File {
		"STDOUT": os.Stdout,
		"STDERR": os.Stderr,
	}
	if out, exists := outputs[output]; !exists {
		return os.Stdout
	} else {
		return out
	}
}

func parseFields(fields map[string]string) []zap.Field {
	fls := make([]zap.Field, 0)
	for k, v := range fields {
		field := zap.String(k, v)
		fls = append(fls, field)
	}
	return fls
}
