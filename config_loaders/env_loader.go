package config_loaders

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"regexp"
	"strings"
)

const (
	logLevelEnv      = "LOG_LEVEL"
	logFieldsEnv     = "LOG_FIELDS"
	logFormatTypeEnv = "LOG_FORMAT_TYPE"
	logOutputType    = "LOG_OUTPUT_TYPE"

	regex = `^\[((\w+\=\w+)+\,?)+\]$`
)

type EnvLoad struct {
	cfg *ConfigLoad
}

func NewEnvLogme() (*EnvLoad, error) {
	lvl, err := logrus.ParseLevel(os.Getenv(logLevelEnv))
	if err != nil {
		return nil, err
	}

	f, _ := regexp.Compile(regex)
	if !f.MatchString(os.Getenv(logFieldsEnv)) {
		return nil, errors.New(fmt.Sprintf("%s environment variable empty", logFieldsEnv))
	}

	s := os.Getenv(logFieldsEnv)
	d := strings.Split(s[1:len(s)-1], ",")
	m := make(map[string]interface{})
	for _, v := range d {
		n := strings.Split(v, "=")
		m[n[0]] = n[1]
	}

	var out io.Writer
	switch os.Getenv(logOutputType) {
	default:
		out = os.Stdout
	}

	var frmt logrus.Formatter
	switch os.Getenv(logFormatTypeEnv) {
	case "json":
		frmt = &logrus.JSONFormatter{}
	default:
		frmt = &logrus.TextFormatter{
			DisableColors:          true,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			FieldMap:               logrus.FieldMap{logrus.FieldKeyTime: "@timestamp"},
		}
	}

	c := &ConfigLoad{
		loaderType: "environment",
		level:      lvl,
		formatter:  frmt,
		fields:     m,
		output:     out,
	}
	m = nil
	return &EnvLoad{cfg: c}, nil
}

func (e *EnvLoad) GetLogLevel() logrus.Level {
	return e.cfg.level
}

func (e *EnvLoad) GetFixedFields() map[string]interface{} {
	return e.cfg.fields
}

func (e *EnvLoad) GetOutputFormatter() logrus.Formatter {
	return e.cfg.formatter
}

func (e *EnvLoad) GetOutput() io.Writer {
	return e.cfg.output
}

func (e *EnvLoad) TypeOf() string {
	return e.cfg.loaderType
}
