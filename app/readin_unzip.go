package app

import (
	"archive/zip"
	"bufio"
	"errors"
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
	var targetFile *zip.File
	foundTarget := false
	for _, file := range zipFile.File {
		// Check if the file is the one we want to process
		if strings.EqualFold(file.Name, dataInputFile) {
			foundTarget = true
			targetFile = file
		}
	}
	if !foundTarget {
		return "", errors.New("target file not found in archive")
	}

	// Open file
	reader, err := targetFile.Open()
	if err != nil {
		return "", err
	}
	defer reader.Close()

	// Read file
	content := ""
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		content += line + "\n"
	}

	return content, nil
}
