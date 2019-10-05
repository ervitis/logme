package main

import (
	"fmt"
	"os"

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
}
