package app

import (
	"os"
)

func WriteFile(outputFile, content string) error {
	err := os.WriteFile(outputFile, []byte(content), 0644)
	return err
}
