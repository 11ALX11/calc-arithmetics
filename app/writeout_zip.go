package app

import (
	"archive/zip"
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

	// Validate dataFile path to prevent creating malicious archives
	if strings.Contains(dataFile, "..") {
		return errors.New("unsafe dataFile path: path traversal detected")
	}
	cleanPath := filepath.Clean(dataFile)
	if filepath.IsAbs(cleanPath) {
		return errors.New("unsafe dataFile path: absolute path detected")
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
