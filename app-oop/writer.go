package app_oop

// Writer represents a type that can write to a file.
type Writer interface {

	// Same as WriteFile() or WriteZipFile() in app package
	WriteFile(outputFile, content string) Writer

	// Setter for dataInputFile attribute
	SetDataInputFile(dataInputFile string) Writer
	// Setter for err attribute
	SetError(err error) Writer
	// Getter for err attribute
	GetError() error
}

// Same as Writer, but with defined fields and setters
type IWriter struct {
	Writer
	err error
}

// Getter for err attribute
func (w IWriter) GetError() error {
	return w.err
}

// Setter for err attribute
func (w IWriter) SetError(err error) Writer {
	w.err = err
	return w
}
