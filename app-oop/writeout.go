package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Writeout represents a type that can write to a txt file.
type Writeout struct {
	err error
}

// Not used with Writeout
func (r *Writeout) SetDataInputFile(dataInputFile string) Writer {
	return r
}

// Getter for err attribute
func (w Writeout) GetError() error {
	return w.err
}

/*
Same as WriteFile() in app package
*/
func (w *Writeout) WriteFile(outputFile, content string) Writer {
	w.err = app.WriteFile(outputFile, content)
	return w
}
