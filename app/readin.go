package app

import (
	"os"
)

/*
Reads a file.

@param inputFile - a filepath to an input file.

@return (string, error) - contents of a file and error if failed to read a file. Nil if success.
*/
func ReadFile(inputFile string) (string, error) {
	content, err := os.ReadFile(inputFile)
	return string(content), err
}
