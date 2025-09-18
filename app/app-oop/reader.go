package app_oop

// Reader represents a type that can read from a file.
type Reader interface {
	// Same as ReadFile() or ReadZipFile() in app package
	ReadFile(inputFile string) (string, error)

	// Setter for dataInputFile attribute
	SetDataInputFile(dataInputFile string) Reader
}
