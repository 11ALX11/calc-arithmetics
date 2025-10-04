package app_oop

// IReaderDecorator is an interface that represents a decorator for a Reader interface.
type IReaderDecorator struct {
	Reader
	wrappee Reader
}

// NewReaderDecorator is a constructor for a decorator.
// func NewReaderDecorator(reader Reader) *IReaderDecorator {
// 	f := new(IReaderDecorator)
// 	f.wrappee = reader
// 	return f
// }

// Getter for content attribute
func (r IReaderDecorator) GetContent() string {
	return r.wrappee.GetContent()
}

// Setter for content attribute
func (r *IReaderDecorator) SetContent(s string) Reader {
	r.wrappee.SetContent(s)
	return r
}

// Getter for err attribute
func (r IReaderDecorator) GetError() error {
	return r.wrappee.GetError()
}

// Setter for content attribute
func (r *IReaderDecorator) SetError(e error) Reader {
	r.wrappee.SetError(e)
	return r
}

// Getter for both content and error.
// Ex: content, err := reader.GetContentError()
func (r IReaderDecorator) GetContentError() (string, error) {
	return r.wrappee.GetContentError()
}

/*
Same as ReadFile() in app package; modify to add decorator functionality
*/
func (r *IReaderDecorator) ReadFile(inputFile string) Reader {

	r.wrappee.ReadFile(inputFile)
	if r.GetError() != nil {
		return r
	}

	// modify here to insert action

	return r
}
