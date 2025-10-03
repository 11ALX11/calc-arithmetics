package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Unzip represents a type that can read a zip file. Implements ReaderDecorator
type Unzip struct {
	wrappee Reader

	dataInputFile string
}

// NewUnzip is a constructor for Unzip.
func NewUnzip(reader Reader, dataInputFile string) Reader {
	u := new(Unzip)
	u.wrappee = reader
	u.dataInputFile = dataInputFile
	return u
}

// Getter for content attribute
func (u Unzip) GetContent() string {
	return u.wrappee.GetContent()
}

// Setter for content attribute
func (u *Unzip) SetContent(s string) Reader {
	u.wrappee.SetContent(s)
	return u
}

// Getter for err attribute
func (u Unzip) GetError() error {
	return u.wrappee.GetError()
}

// Setter for content attribute
func (u *Unzip) SetError(e error) Reader {
	u.wrappee.SetError(e)
	return u
}

// Getter for both content and error.
// Ex: content, err := reader.GetContentError()
func (u Unzip) GetContentError() (string, error) {
	return u.wrappee.GetContentError()
}

/*
Same as ReadZipFile() in app package but
reads a zip archive from a GetContent() string
to make possible decorator chains.
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
