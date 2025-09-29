package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Writer represents a type that can write to a txt file.
type WriteoutZip struct {
	dataFileInArchive string

	err error
}

// Setter for dataInputFile attribute
func (w *WriteoutZip) SetDataInputFile(dataFileInArchive string) Writer {
	w.dataFileInArchive = dataFileInArchive
	return w
}

// Getter for err attribute
func (w WriteoutZip) GetError() error {
	return w.err
}

/*
Same as WriteZipFile() in app package
*/
func (w *WriteoutZip) WriteFile(outputFile, content string) Writer {
	w.err = app.WriteZipFile(outputFile, content, w.dataFileInArchive)
	return w
}
