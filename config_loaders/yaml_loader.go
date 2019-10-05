package config_loaders

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type YamlLoad struct {
	cfg *ConfigLoad
}

/*
logger:
	level:
	fields:
		- field1
		- field2
	format:
		type: json|txt
	output:
		type:
 */
func parseYaml(pathYaml string) ([]byte, error) {
	type y struct {
		L string `yaml:"level"`
		F []string `yaml:"fields"`
		Fo struct {
			T string `yaml:"type"`
		} `yaml:"format"`
		O struct {
			T string `yaml:"type"`
		} `yaml:"output"`
	}

	data := y{}

	b, err := ioutil.ReadFile(pathYaml)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return yaml.Marshal(data)
}

func NewYamlLogme(pathYaml string) (*YamlLoad, error) {
	var data map[string]interface{}

	b, err := parseYaml(pathYaml)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	lvl, err := logrus.ParseLevel(data["level"].(string))
	if err != nil {
		return nil, err
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
