package app_oop

// ReaderDecorator is an interface that represents a decorator for a Reader interface.
type WriterDecorator interface {
	Writer
}

// Same as ReaderDecorator, but with defined fields, setters and getters
type IWriterDecorator struct {
	WriterDecorator
	wrappee Writer
}

// NewWriterDecorator is a constructor for a decorator.
// func NewWriterDecorator(reader Writer) *IWriterDecorator {
// 	w := new(IWriterDecorator)
// 	w.wrappee = reader
// 	return w
// }

// Getter for err attribute
func (w IWriterDecorator) GetError() error {
	return w.wrappee.GetError()
}

// Setter for content attribute
func (w *IWriterDecorator) SetError(e error) Writer {
	w.wrappee.SetError(e)
	return w
}
