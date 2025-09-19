package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// ReadinUnzip represents a type that can read a zip file.
type ReadinUnzip struct {
	dataInputFile string

	content string
	err     error
}

// Accepts ReaderVisitor implementations and calls DoForReadinUnzip()
func (r *ReadinUnzip) Accept(v ReaderVisitor) {
	v.DoForReadinUnzip(r)
}

// Setter for dataInputFile attribute
func (r *ReadinUnzip) SetDataInputFile(dataInputFile string) Reader {
	r.dataInputFile = dataInputFile
	return r
}

// Getter for content attribute
func (r ReadinUnzip) GetContent() string {
	return r.content
}

// Setter for content attribute
func (r *ReadinUnzip) SetContent(s string) Reader {
	r.content = s
	return r
}

// Getter for err attribute
func (r ReadinUnzip) GetError() error {
	return r.err
}

// Setter for content attribute
func (r *ReadinUnzip) SetError(e error) Reader {
	r.err = e
	return r
}

// Getter for both content and error.
// Ex: content, err := reader.GetContentError()
func (r ReadinUnzip) GetContentError() (string, error) {
	return r.content, r.err
}

/*
Same as ReadZipFile() in app package
*/
func (r *ReadinUnzip) ReadFile(inputFile string) Reader {
	r.content, r.err = app.ReadZipFile(inputFile, r.dataInputFile)
	return r
}
