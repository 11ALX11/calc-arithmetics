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

// // Setter for dataInputFile attribute
// func (r *ReaderDecorator) SetDataInputFile(dataInputFile string) Reader {
// 	return r.wrappee.SetDataInputFile(dataInputFile)
// }

// // Getter for content attribute
// func (r ReaderDecorator) GetContent() string {
// 	return r.wrappee.GetContent()
// }

// // Setter for content attribute
// func (r *ReaderDecorator) SetContent(s string) Reader {
// 	return r.wrappee.SetContent(s)
// }

// // Getter for err attribute
// func (r ReaderDecorator) GetError() error {
// 	return r.wrappee.GetError()
// }

// // Setter for content attribute
// func (r *ReaderDecorator) SetError(e error) Reader {
// 	return r.wrappee.SetError(e)
// }

// // Getter for both content and error.
// // Ex: content, err := reader.GetContentError()
// func (r ReaderDecorator) GetContentError() (string, error) {
// 	return r.wrappee.GetContentError()
// }

// /*
// Same as ReadFile() in app package
// */
// func (r *ReaderDecorator) ReadFile(inputFile string) Reader {
// 	// modify here to insert action
// 	return r.wrappee.ReadFile(inputFile)
// }
