package golanglog

import (
	"github.com/ervitis/logme/v2/common"
	"github.com/ervitis/logme/v2/config_loaders"
	"log"
)

type (
	WrapperLog struct {
		Log *log.Logger
	}
)

func parseFlags(cfg *config_loaders.LoaderModel) int {
	flags := cfg.InitialFields["flags"]
	if flags == "" {
		return log.LstdFlags
	}
	mFlags := map[string]int{
		"LDATE":         log.Ldate,
		"LTIME":         log.Ltime,
		"LMICROSECONDS": log.Lmicroseconds,
		"LLONGFILE":     log.Llongfile,
		"LSHORTFILE":    log.Lshortfile,
		"LUTC":          log.LUTC,
		"LMSGPREFIX":    log.Lmsgprefix,
	}
	if f, ok := mFlags[flags]; !ok {
		return log.LstdFlags
	} else {
		return f
	}
}

func NewLog(configModel *config_loaders.LoaderModel) (*WrapperLog, error) {
	prefix := configModel.InitialFields["prefix"]

	return &WrapperLog{
		Log: log.New(
			common.ParseOutput(configModel.Output),
			prefix,
			parseFlags(configModel),
		)}, nil
}

func (l *WrapperLog) L() *WrapperLog {
	return l
}

func (l *WrapperLog) Debug(msg string) {
	l.Log.Print(msg)
}

func (l *WrapperLog) Info(msg string) {
	l.Log.Print(msg)
}

func (l *WrapperLog) Warn(msg string) {
	l.Log.Print(msg)
}

func (l *WrapperLog) Error(msg string) {
	l.Log.Print(msg)
}

func (l *WrapperLog) Panic(msg string) {
	l.Log.Panic(msg)
}
