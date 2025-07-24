package app

import (
	"os"
)

func ReadFile(inputFile string) (string, error) {
	content, err := os.ReadFile(inputFile)
	return string(content), err
}
