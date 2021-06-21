package config_loaders

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	YamlLoader struct {
		model *LoaderModel
		Path  string
	}
)

func NewYamlLoader() (*YamlLoader, error) {
	return &YamlLoader{}, nil
}

func (e *YamlLoader) GetPath() string {
	return e.Path
}

func (e *YamlLoader) SetPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("yaml loader error set path: %w", err)
	}
	e.Path = path
	return nil
}

func (e *YamlLoader) Marshal(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}

func (e *YamlLoader) Unmarshal(in []byte, out interface{}) error {
	return yaml.Unmarshal(in, out)
}
