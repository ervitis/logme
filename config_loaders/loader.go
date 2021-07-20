package config_loaders

import (
	"fmt"
	"io/ioutil"
)

type (
	LoaderModel struct {
		Level         string            `json:"level" yaml:"level"`
		Encoding      string            `json:"encoding" yaml:"encoding"`
		Output        string            `json:"output,omitempty" yaml:"output,omitempty"`
		InitialFields map[string]string `json:"initialFields,omitempty" yaml:"initialFields,omitempty"`
	}

	Loader interface {
		Marshal(interface{}) ([]byte, error)
		Unmarshal([]byte, interface{}) error
		SetPath(string) error
		GetPath() string
	}

	LoggermeConfig struct {
		Path   string
		loader Loader
	}
)

func NewLoggermeConfig(loader Loader) *LoggermeConfig {
	return &LoggermeConfig{loader: loader}
}

func (lc *LoggermeConfig) OpenFile(path string) ([]byte, error) {
	if path == "" {
		return nil, nil
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("open config file error: %w", err)
	}
	return data, nil
}

func (lc *LoggermeConfig) LoggermeConfigLoader() (*LoaderModel, error) {
	data, err := lc.OpenFile(lc.loader.GetPath())
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	var model LoaderModel
	if err := lc.loader.Unmarshal(data, &model); err != nil {
		return nil, fmt.Errorf("error unmarshaling data: %w", err)
	}
	return &model, nil
}
