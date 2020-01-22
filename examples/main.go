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

func func1(l logme.Loggerme) {
	func2(l)
}

func func2(l logme.Loggerme) {
	func3(l)
}

func func3(log logme.Loggerme) {
	log.L().WithError(errors.New("ups")).Errorf("i am inside a function")
}

func main() {
	_ = os.Setenv("LOG_LEVEL", "DEBUG")
	_ = os.Setenv("LOG_FIELDS", "[component=example,service=example]")
	cfg, err := config_loaders.NewEnvLogme()
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.TypeOf())

	log := logme.NewLogme(cfg)
	log.L().WithField("traceID", "111111111").Debug("hello world")

	path, err := filepath.Abs("examples/log_config.yaml")
	if err != nil {
		panic(err)
	}

	cfg2, err := config_loaders.NewYamlLogme(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg2.TypeOf())
	log2 := logme.NewLogme(cfg2).L()

	log2.Info("hello world 2")

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

	log2.Debug("hello world 3")
	log2.Error("oh no...")
	log.L().Info("hi!!")
	log.L().WithField("traceId", "9999999999").WithError(errors.New("ups")).Errorf("something is broken")

	func1(log)
}
