package config_loaders

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
)

type YamlLoad struct {
	cfg *ConfigLoad
}

type logYaml struct {
	Level string `yaml:"level"`
	Fields map[string]interface{} `yaml:"fields"`
	Format struct {
		Type string `yaml:"type"`
	} `yaml:"format"`
	Output struct {
		Type string `yaml:"type"`
	} `yaml:"output"`
}

func parseYaml(pathYaml string) (*logYaml, error) {
	data := logYaml{}

	b, err := ioutil.ReadFile(pathYaml)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func NewYamlLogme(pathYaml string) (*YamlLoad, error) {
	data, err := parseYaml(pathYaml)
	if err != nil {
		return nil, err
	}

	lvl, err := logrus.ParseLevel(data.Level)
	if err != nil {
		return nil, err
	}

	var out io.Writer
	switch data.Output.Type {
	default:
		out = os.Stdout
	}

	var frmt logrus.Formatter
	switch data.Format.Type {
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
		loaderType: "yaml",
		level:      lvl,
		formatter:  frmt,
		fields:     data.Fields,
		output:     out,
	}
	return &YamlLoad{cfg: c}, nil
}

func (e *YamlLoad) GetLogLevel() logrus.Level {
	return e.cfg.level
}

func (e *YamlLoad) GetFixedFields() map[string]interface{} {
	return e.cfg.fields
}

func (e *YamlLoad) GetOutputFormatter() logrus.Formatter {
	return e.cfg.formatter
}

func (e *YamlLoad) GetOutput() io.Writer {
	return e.cfg.output
}

func (e *YamlLoad) TypeOf() string {
	return e.cfg.loaderType
}
