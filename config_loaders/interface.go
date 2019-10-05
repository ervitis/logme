package config_loaders

import (
	"github.com/sirupsen/logrus"
	"io"
)

type ConfigLoad struct {
	loaderType string
	fields     map[string]interface{}
	level      logrus.Level
	formatter  logrus.Formatter
	output     io.Writer
}

type ConfigLoader interface {
	GetLogLevel() logrus.Level
	GetFixedFields() map[string]interface{}
	GetOutputFormatter() logrus.Formatter
	TypeOf() string
	GetOutput() io.Writer
}

func (e *ConfigLoad) GetLogLevel() logrus.Level {
	return e.level
}

func (e *ConfigLoad) GetFixedFields() map[string]interface{} {
	return e.fields
}

func (e *ConfigLoad) GetOutputFormatter() logrus.Formatter {
	return e.formatter
}

func (e *ConfigLoad) GetOutput() io.Writer {
	return e.output
}

func (e *ConfigLoad) TypeOf() string {
	return e.loaderType
}
