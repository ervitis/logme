package config_loaders

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

type JsonLoad struct {
	ConfigLoader
}

type logJson struct {
	Level string `json:"level"`
	Fields map[string]interface{} `json:"fields"`
	Format struct {
		Type string `json:"type"`
	} `json:"format"`
	Output struct {
		Type string `json:"type"`
	} `json:"output"`
}

func parseJson(pathJson string) (*logJson, error) {
	data := logJson{}

	b, err := ioutil.ReadFile(pathJson)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func NewJsonLogme(pathJson string) (*JsonLoad, error) {
	data, err := parseJson(pathJson)
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

	frmt := CommonFormatter(data.Format.Type)

	c := &ConfigLoad{
		loaderType: "json",
		level:      lvl,
		formatter:  frmt,
		fields:     data.Fields,
		output:     out,
	}
	return &JsonLoad{ c}, nil
}
