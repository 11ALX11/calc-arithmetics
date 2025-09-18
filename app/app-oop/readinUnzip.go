package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// ReadinUnzip represents a type that can read a zip file.
type ReadinUnzip struct {
	dataInputFile string
}

// Setter for dataInputFile attribute
func (r ReadinUnzip) SetDataInputFile(dataInputFile string) Reader {
	r.dataInputFile = dataInputFile
	return r
}

/*
Same as ReadZipFile() in app package
*/
func (r ReadinUnzip) ReadFile(inputFile string) (string, error) {
	return app.ReadZipFile(inputFile, r.dataInputFile)
}
