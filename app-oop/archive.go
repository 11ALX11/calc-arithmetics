package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Writer represents a type that can write to a txt file.
type Archive struct {
	IWriterDecorator
	dataFileInArchive string
}

// NewDecrypt is a constructor for Decrypt.
func NewArchive(writer Writer, dataFileInArchive string) Writer {
	a := new(Archive)
	a.wrappee = writer
	a.dataFileInArchive = dataFileInArchive
	return a
}

// Setter for dataInputFile attribute
func (a *Archive) SetDataInputFile(dataFileInArchive string) Writer {
	a.dataFileInArchive = dataFileInArchive
	return a
}

/*
Uses GetZipData() instead of GetZipData() (in app package)
to modify content to an archive to make possible decorator chains.

Skips writing if caught an error.
*/
func (a *Archive) WriteFile(outputFile, content string) Writer {

	mod_content, err := app.GetZipData(content, a.dataFileInArchive)
	if err != nil {
		a.SetError(err)
		return a
	}

	a.wrappee.WriteFile(outputFile, mod_content)
	return a
}
