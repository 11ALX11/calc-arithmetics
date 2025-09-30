package app_oop

// ReaderDecorator represents a decorator for a Reader interface.
type ReaderDecorator interface {
	Reader
}

// NewReaderDecorator is a constructor for a decorator.
// func NewReaderDecorator(reader Reader) *ReaderDecorator {
// 	f := new(ReaderDecorator)
// 	f.wrappee = reader
// 	return f
// }
