package common

import (
	"os"
	"strings"
)

func ParseOutput(output string) *os.File {
	output = strings.ToUpper(output)
	outputs := map[string]*os.File{
		"STDOUT": os.Stdout,
		"STDERR": os.Stderr,
	}
	if out, exists := outputs[output]; !exists {
		return os.Stdout
	} else {
		return out
	}
}
