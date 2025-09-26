package app_oop

// Writer represents a type that can write to a file.
type Writer interface {

	// Same as WriteFile() or WriteZipFile() in app package
	WriteFile(outputFile, content string) Writer

	// Setter for dataInputFile attribute
	SetDataInputFile(dataInputFile string) Writer
	// Getter for err attribute
	GetError() error
}
