package config_loaders

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	JsonLoader struct {
		model *LoaderModel
		Path  string
	}
)

func NewJsonLoader() (*JsonLoader, error) {
	return &JsonLoader{}, nil
}

func (e *JsonLoader) GetPath() string {
	return e.Path
}

func (e *JsonLoader) SetPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("json loader error set path: %w", err)
	}
	e.Path = path
	return nil
}

func (e *JsonLoader) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (e *JsonLoader) Unmarshal(in []byte, out interface{}) error {
	return json.Unmarshal(in, out)
}
