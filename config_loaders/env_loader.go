package config_loaders

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type (
	EnvLoader struct {
		model *LoaderModel
	}
)

func NewEnvLoader() (*EnvLoader, error) {
	loader := &EnvLoader{}
	if err := loader.Init(); err != nil {
		return nil, fmt.Errorf("error loading env variables %w", err)
	}

	return &EnvLoader{
		model: loader.model,
	}, nil
}

var (
	LogLevelEmptyErr = "%s is empty"
)

const (
	LogLevel            = "LOG_LEVEL"
	LogEncoding         = "LOG_ENCODING"
	LogOutputpath      = "LOG_OUTPUTPATH"
	LogInitialfields    = "LOG_INITIALFIELDS"
)

func (e *EnvLoader) Init() error {
	v := os.Getenv(LogLevel)
	if v == "" {
		return fmt.Errorf(LogLevelEmptyErr, LogLevel)
	}
	e.model.Level = v

	v = os.Getenv(LogEncoding)
	if v == "" {
		return fmt.Errorf(LogLevelEmptyErr, LogEncoding)
	}
	e.model.Encoding = v

	v = os.Getenv(LogOutputpath)
	if v == "" {
		return fmt.Errorf(LogLevelEmptyErr, LogOutputpath)
	}
	e.model.Output = v

	v = os.Getenv(LogInitialfields)
	if v == "" {
		return fmt.Errorf(LogLevelEmptyErr, LogInitialfields)
	}
	iflds := make(map[string]string)
	for _, m := range strings.Split(v, ",") {
		for _, n := range strings.Split(m, "=") {
			iflds[n] = n
		}
	}
	if len(iflds) > 0 {
		e.model.InitialFields = iflds
	}

	return nil
}

func (e *EnvLoader) GetPath() string {
	return ""
}

func (e *EnvLoader) SetPath(_ string) error {
	return nil
}

func (e *EnvLoader) Marshal(_ interface{}) ([]byte, error) {
	return json.Marshal(e.model)
}

func (e *EnvLoader) Unmarshal(_ []byte, out interface{}) error {
	out = e.model
	return nil
}
