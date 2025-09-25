package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Writer represents a type that can write to a txt file.
type Writein struct {
	err error
}

// Getter for err attribute
func (w Writein) GetError() error {
	return w.err
}

/*
Same as WriteFile() in app package
*/
func (w *Writein) WriteFile(outputFile, content string) Writer {
	w.err = app.WriteFile(outputFile, content)
	return w
}
