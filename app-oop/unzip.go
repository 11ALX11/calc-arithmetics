package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Unzip represents a type that can read a zip file. Implements ReaderDecorator
type Unzip struct {
	IReaderDecorator
	dataInputFile string
}

// NewUnzip is a constructor for Unzip decorator.
func NewUnzip(reader Reader, dataInputFile string) Reader {
	u := new(Unzip)
	u.wrappee = reader
	u.dataInputFile = dataInputFile
	return u
}

/*
Uses ReadZipData() instead of ReadZipFile() (in app package)
to read a zip archive from a GetContent() string
to make possible decorator chains.

Skips action if reader already has an error.
*/
func (u *Unzip) ReadFile(inputFile string) Reader {

	u.wrappee.ReadFile(inputFile)
	if u.GetError() != nil {
		return u
	}

	content, err := app.ReadZipData(u.GetContent(), u.dataInputFile)
	u.SetContent(content)
	u.SetError(err)

	return u
}
