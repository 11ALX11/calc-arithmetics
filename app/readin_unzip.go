package app

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const DataFileInArchive = "data.txt" // A file inside of an archive to extract or write contents to.

type ZipFileReader interface {
	Files() []*zip.File
}

// Wrapper for *zip.Reader
type ZipReaderWrapper struct {
	*zip.Reader
}

func (z *ZipReaderWrapper) Files() []*zip.File {
	return z.Reader.File
}

// Wrapper for *zip.ReadCloser
type ZipReadCloserWrapper struct {
	*zip.ReadCloser
}

func (z *ZipReadCloserWrapper) Files() []*zip.File {
	return z.ReadCloser.File
}

/*
ReadZipFile reads a zip archive and returns contents of an dataInputFile.
If dataInputFile is not found, it falls back to the first file in the archive.

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

	return readZipWithReader(&ZipReadCloserWrapper{zipFile}, dataInputFile)
}

func readZipWithReader(zipFile ZipFileReader, dataInputFile string) (string, error) {
	// Iterate through each file in the archive
	var targetFile *zip.File = nil
	foundTarget := false
	for i, file := range zipFile.Files() {
		// Check if the file is the one we want to process
		if strings.EqualFold(file.Name, dataInputFile) {
			foundTarget = true
			targetFile = file
			break
		} else if i == 0 {
			targetFile = file // remember first file in archive in case none match dataInputFile
		}
	}
	if targetFile == nil {
		return "", fmt.Errorf("file not found in archive %q", dataInputFile)
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

/*
This is a modification to a ReadZipFile().

ReadZipData reads a zip archive from a string and returns contents of a dataInputFile.
If dataInputFile is not found, it falls back to the first file in the archive.

@param zipData - binary data of an archive

@param dataInputFile - a file inside of an archive to extract contents of. Usually used const DataFileInArchive

@return (string, error) - content of a file in an archive and any error that can occure while reading zip. Nil if no error happened.
*/
func ReadZipData(zipData, dataInputFile string) (string, error) {
	// Convert the string to a byte slice
	data := []byte(zipData)

	// Create a bytes.Reader from the byte slice
	reader := bytes.NewReader(data)

	// Create a zip.Reader from the bytes.Reader
	zipReader, err := zip.NewReader(reader, int64(len(data)))
	if err != nil {
		return "", err
	}

	return readZipWithReader(&ZipReaderWrapper{zipReader}, dataInputFile)
}
