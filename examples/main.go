package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ervitis/logme/hooks"

	"github.com/ervitis/logme"
	"github.com/ervitis/logme/config_loaders"
)

func main() {
	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	_ = os.Setenv("LOG_FIELDS", "[component=example,service=example]")
	cfg, err := config_loaders.NewEnvLogme()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.TypeOf())

	log := logme.NewLogme(cfg)
	log.Debug("hello world", logme.Metadata{TraceID: "12345"})

	path, err := filepath.Abs("examples/log_config.yaml")
	if err != nil {
		panic(err)
	}

	cfg2, err := config_loaders.NewYamlLogme(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg2.TypeOf())
	log = logme.NewLogme(cfg2)

	log.Info("hello world 2")

	path, err = filepath.Abs("examples/log_config.json")
	if err != nil {
		panic(err)
	}

	cfg3, err := config_loaders.NewJsonLogme(path)
	if err != nil {
		panic(err)
	}

	myHooks := []logme.Hook{
		hooks.NewStackHook(),
	}
	log = logme.NewLogme(cfg3, myHooks...)

	log.Debug("hello world 3")
	log.Error("oh no...").WithError(errors.New("ups"))
}
