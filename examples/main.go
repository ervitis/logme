package main

import (
	"github.com/ervitis/logme/v2"
)

func main() {
	log := logme.NewLogme()

	log.L().Info("Hi")
}
