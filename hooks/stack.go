package hooks

import (
	"github.com/facebookgo/stack"
	"github.com/sirupsen/logrus"
	"strings"
)

type (
	StackHook struct {
		CallerLevels []logrus.Level
		StackLevels  []logrus.Level
	}
)

const (
	stackJump          = 4
	fieldLessStackJump = 6
)

func NewStackHook() *StackHook {
	return &StackHook{
		CallerLevels: logrus.AllLevels,
		StackLevels:  []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel},
	}
}

func (h *StackHook) Levels() []logrus.Level {
	return h.CallerLevels
}

func (h *StackHook) Fire(e *logrus.Entry) error {
	skipFrames := stackJump
	if len(e.Data) == 0 {
		skipFrames = fieldLessStackJump
	}

	var st stack.Stack
	frames := stack.Callers(skipFrames)

	for _, frame := range frames {
		if !strings.Contains(frame.File, "sirupsen/logrus") && !strings.Contains(frame.File,"hooks/stack.go") {
			st = append(st, frame)
		}
	}

	if len(st) > 0 {
		e.Data["function"] = st[0].String()

		for _, level := range h.StackLevels {
			if e.Level == level {
				e.Data["errTrace"] = st
				break
			} else {
				delete(e.Data, "errTrace")
			}
		}
	}
	return nil
}
