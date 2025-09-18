package app_oop

// Reader represents a type that can read from a file.
type Reader interface {

	// Accepts ReaderVisitor implementations and calls doFor*Something*()
	Accept(v ReaderVisitor)

	// Same as ReadFile() or ReadZipFile() in app package
	ReadFile(inputFile string) Reader

	// Setter for dataInputFile attribute
	SetDataInputFile(dataInputFile string) Reader
	// Getter for content attribute
	GetContent() string
	// Setter for content attribute
	SetContent(content string) Reader
	// Getter for err attribute
	GetError() error
	// Setter for err attribute
	SetError(err error) Reader
	// Getter for both content and error.
	// Ex: content, err := reader.GetContentError()
	GetContentError() (string, error)
}
