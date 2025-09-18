package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Readin represents a type that can read a text file.
type Readin struct {
}

// Not used with Readin
func (r Readin) SetDataInputFile(dataInputFile string) Reader {
	return r
}

/*
Same as ReadFile() in app package
*/
func (r Readin) ReadFile(inputFile string) (string, error) {
	return app.ReadFile(inputFile)
}
