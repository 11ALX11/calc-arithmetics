package app

import (
	"os"
)

/*
Writes to a file.

@param outputFile - a filepath to a output file.

@param content - a content to write.

@return error - error if failed to write to a file. Nil if success.
*/
func WriteFile(outputFile, content string) error {
	err := os.WriteFile(outputFile, []byte(content), 0644)
	return err
}
