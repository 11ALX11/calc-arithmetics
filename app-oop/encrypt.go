package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Encrypt is a decorator that encrypts content before writing.
type Encrypt struct {
	IWriterDecorator
	keyPath string
}

// NewEncrypt is a constructor for Encrypt decorator.
func NewEncrypt(writer Writer, keyPath string) Writer {
	return &Encrypt{
		IWriterDecorator{wrappee: writer},
		keyPath,
	}
}

// Setter for keyPath attribute
func (e *Encrypt) SetKeyPath(keyPath string) Writer {
	e.keyPath = keyPath
	return e
}

/*
Encrypts content.
Skips writing if caught an error.
*/
func (e *Encrypt) WriteFile(outputFile, content string) Writer {

	mod_content, err := app.EncryptFileKey(content, e.keyPath)
	if err != nil {
		e.SetError(err)
		return e
	}

	e.wrappee.WriteFile(outputFile, mod_content)
	return e
}
