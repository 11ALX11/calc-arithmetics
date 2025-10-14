package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Readin represents a type that can read a text file.
type Readin struct {
	content string
	err     error
}

// NewReadin is a constructor for Readin.
func NewReadin() Reader {
	return &Readin{}
}

// Getter for content attribute
func (r Readin) GetContent() string {
	return r.content
}

// Setter for content attribute
func (r *Readin) SetContent(s string) Reader {
	r.content = s
	return r
}

// Getter for err attribute
func (r Readin) GetError() error {
	return r.err
}

// Setter for err attribute
func (r *Readin) SetError(e error) Reader {
	r.err = e
	return r
}

// Getter for both content and error.
// Ex: content, err := reader.GetContentError()
func (r Readin) GetContentError() (string, error) {
	return r.content, r.err
}

/*
Same as ReadFile() in app package
*/
func (r *Readin) ReadFile(inputFile string) Reader {
	r.content, r.err = app.ReadFile(inputFile)
	return r
}
