package main

import (
	"fmt"
	"os"
	"path/filepath"

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
	log.Debug("hello world", "12345")

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

	log.Info("hello world 2", "3420702")

	path, err = filepath.Abs("examples/log_config.json")
	if err != nil {
		panic(err)
	}

	cfg3, err := config_loaders.NewJsonLogme(path)
	if err != nil {
		panic(err)
	}

	log = logme.NewLogme(cfg3)

	log.Debug("hello world 3", "5353462546")
}
