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
