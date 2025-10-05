package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Writeout represents a type that can write to a txt file.
type Writeout struct {
	IWriter
}

// NewWriteout is a constructor for Writeout.
func NewWriteout() Writer {
	w := new(Writeout)
	return w
}

/*
Same as WriteFile() in app package
*/
func (w *Writeout) WriteFile(outputFile, content string) Writer {
	w.err = app.WriteFile(outputFile, content)
	return w
}
