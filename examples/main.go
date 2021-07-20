package main

import (
	"github.com/ervitis/logme/v2"
	"github.com/ervitis/logme/v2/config_loaders"
	"path/filepath"
)

func main() {
	yamlLoader, err := config_loaders.NewYamlLoader()
	if err != nil {
		panic(err)
	}
	yamlFile, err := filepath.Abs("examples/log_config.yaml")
	if err != nil {
		panic(err)
	}

	_ = yamlLoader.SetPath(yamlFile)

	log, err := logme.NewLogme(logme.WithConfigLoader(yamlLoader))
	if err != nil {
		panic(err)
	}

	log.L().Debug("hello")
}
