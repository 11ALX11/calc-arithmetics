package app_oop

// WriterDecorator is an interface that represents a decorator for a Writer interface.
// Applied first-inner to last-outer
// or last decorator to apply will be executed first
type WriterDecorator interface {
	Writer
}

// Same as WriterDecorator, but with defined fields, setters and getters
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
