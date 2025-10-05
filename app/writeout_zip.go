package app

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/*
WriteZipFile creates a zip file containing a text file with the specified content.

@param outputFile - a filepath to a output file.

@param content - a content to write.

@param dataFile - a file inside of an archive to write contents to. Usually used const DataFileInArchive

@return error - error if failed to write/create a zip. Nil if success.
*/
func WriteZipFile(outputFile, content, dataFile string) (err error) {

	if err := validateDataFile(dataFile); err != nil {
		return err
	}

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "ziptemp")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir) // Clean up the temp directory

	// Create the dataFile with the provided content
	dataFilePath := filepath.Join(tempDir, dataFile)
	if err := WriteFile(dataFilePath, content); err != nil {
		return err
	}

	// Create the zip file
	zipFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := zipFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if cerr := zipWriter.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	// Add dataFile to the zip file
	if err := writeFileToZip(zipWriter, dataFilePath, dataFile); err != nil {
		return err
	}

	return nil
}

func writeFileToZip(w *zip.Writer, file, fileInZip string) error {
	src, err := os.Open(file)
	if err != nil {
		return err
	}

	f, err := w.Create(fileInZip)
	if err != nil {
		return err
	}
	if _, err := io.Copy(f, src); err != nil {
		return err
	}

	return nil
}

func validateDataFile(dataFilePath string) error {
	// Validate dataFile path to prevent creating malicious archives
	if strings.Contains(dataFilePath, "..") {
		return errors.New("unsafe dataFile path: path traversal detected")
	}
	cleanPath := filepath.Clean(dataFilePath)
	if filepath.IsAbs(cleanPath) {
		return errors.New("unsafe dataFile path: absolute path detected")
	}
	return nil
}

/*
GetZipData returns a binary string representing a zip file containing a text file with the specified content.

@param content - a content to archive.

@param dataFile - a file inside of an archive to write contents to. Usually used const DataFileInArchive

@return (string, error) - string with archived content, error if failed to archive. Nil if success.
*/
func GetZipData(content, dataFile string) (s string, err error) {

	if err := validateDataFile(dataFile); err != nil {
		return "", err
	}

	// Create a buffer to hold the zip data
	var zipBuffer *bytes.Buffer = new(bytes.Buffer)
	var zipWriter *zip.Writer = zip.NewWriter(zipBuffer)

	f, err := zipWriter.Create(dataFile)
	if err != nil {
		return "", err
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		return "", err
	}

	// Close zipWriter
	if err := zipWriter.Close(); err != nil {
		return "", err
	}

	return zipBuffer.String(), nil
}
