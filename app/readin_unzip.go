package app

import (
	"archive/zip"
	"fmt"
	"io"
	"strings"
)

const DataFileInArchive = "data.txt" // A file inside of an archive to extract or write to contents.

/*
ReadZipFile reads a zip archive and returns contents of an dataInputFile.

@param inputArchive - filepath to an archive

@param dataInputFile - a file inside of an archive to extract contents of. Usually used const DataFileInArchive

@return (string, error) - content of a file in an archive and any error that can occure while reading zip. Nil if no error happened.
*/
func ReadZipFile(inputArchive, dataInputFile string) (string, error) {
	// Open the ZIP archive for reading
	zipFile, err := zip.OpenReader(inputArchive)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	// Iterate through each file in the archive
	var targetFile *zip.File = nil
	foundTarget := false
	i := 0
	for _, file := range zipFile.File {
		// Check if the file is the one we want to process
		if strings.EqualFold(file.Name, dataInputFile) {
			foundTarget = true
			targetFile = file
			break
		} else if i == 0 {
			targetFile = file // remember first file in archive in case none match dataInputFile
		}
		i++
	}
	if targetFile == nil {
		return "", fmt.Errorf("file not found in archive %q", inputArchive)
	}
	if !foundTarget {
		// ToDo: info log
	}

	// Open file
	reader, err := targetFile.Open()
	if err != nil {
		return "", err
	}
	defer reader.Close()

	// Read entire file
	buf, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
