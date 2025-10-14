package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Decrypt is a type that decrypts a ciphertext string using AES. Implements ReaderDecorator
type Decrypt struct {
	IReaderDecorator
	keyPath string
}

// NewDecrypt is a constructor for Decrypt decorator.
func NewDecrypt(reader Reader, keyPath string) Reader {
	return &Decrypt{
		IReaderDecorator{wrappee: reader},
		keyPath,
	}
}

/*
Decrypts ciphertext from GetContent().
Skips action if reader already has an error.
*/
func (d *Decrypt) ReadFile(inputFile string) Reader {

	d.wrappee.ReadFile(inputFile)
	if d.GetError() != nil {
		return d
	}

	plaintext, err := d.DecryptFileKey(d.GetContent(), d.keyPath)
	d.SetContent(plaintext)
	d.SetError(err)

	return d
}

// Setter for keyPath attribute
func (d *Decrypt) SetKeyPath(keyPath string) *Decrypt {
	d.keyPath = keyPath
	return d
}

// Same as DecryptFileKey() in app package
func (d Decrypt) DecryptFileKey(ciphertext, keyPath string) (string, error) {
	return app.DecryptFileKey(ciphertext, keyPath)
}

// Same as Decrypt() in app package
func (d Decrypt) Decrypt(ciphertext, key string) (string, error) {
	return app.Decrypt(ciphertext, key)
}
