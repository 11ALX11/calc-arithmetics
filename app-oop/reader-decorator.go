package app_oop

// ReaderDecorator is an interface that represents a decorator for a Reader interface.
// Applied first-first to last-last
type ReaderDecorator interface {
	Reader
}

// Same as ReaderDecorator, but with defined fields, setters and getters
type IReaderDecorator struct {
	ReaderDecorator
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
