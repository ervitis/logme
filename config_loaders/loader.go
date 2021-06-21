package config_loaders

import (
	"fmt"
	"io/ioutil"
)

type (
	LoaderModel struct {
		Level            string            `json:"level" yaml:"level"`
		Encoding         string            `json:"encoding" yaml:"encoding"`
		OutputPaths      []string          `json:"outputPaths,omitempty" yaml:"outputPaths"`
		ErrorOutputPaths []string          `json:"errorOutputPaths" yaml:"errorOutputPaths"`
		InitialFields    map[string]string `json:"initialFields" yaml:"initialFields"`
	}

	Loader interface {
		Marshal(interface{}) ([]byte, error)
		Unmarshal([]byte, interface{}) error
		SetPath(string) error
		GetPath() string
	}

	LoggermeConfig struct {
		Path string
	}
)

func (lc *LoggermeConfig) OpenFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open config file error: %w", err)
	}
	return data, nil
}

func (lc *LoggermeConfig) LoggermeConfigLoader(loader Loader) (*LoaderModel, error) {
	data, err := lc.OpenFile(loader.GetPath())
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	var model *LoaderModel
	if err := loader.Unmarshal(data, model); err != nil {
		return nil, fmt.Errorf("error unmarshaling data: %w", err)
	}
	return model, nil
}
