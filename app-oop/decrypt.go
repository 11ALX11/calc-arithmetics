package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Decrypt is a type that decrypts a ciphertext string using AES. Implements ReaderDecorator
type Decrypt struct {
	wrappee Reader

	keyPath string
}

// NewDecrypt is a constructor for Decrypt.
func NewDecrypt(reader Reader, keyPath string) Reader {
	d := new(Decrypt)
	d.wrappee = reader
	d.keyPath = keyPath
	return d
}

// Getter for content attribute
func (d Decrypt) GetContent() string {
	return d.wrappee.GetContent()
}

// Setter for content attribute
func (d *Decrypt) SetContent(s string) Reader {
	d.wrappee.SetContent(s)
	return d
}

// Getter for err attribute
func (d Decrypt) GetError() error {
	return d.wrappee.GetError()
}

// Setter for content attribute
func (d *Decrypt) SetError(e error) Reader {
	d.wrappee.SetError(e)
	return d
}

// Getter for both content and error.
// Ex: content, err := reader.GetContentError()
func (d Decrypt) GetContentError() (string, error) {
	return d.wrappee.GetContentError()
}

/*
Same as ReadFile() in app package
*/
func (d *Decrypt) ReadFile(inputFile string) Reader {

	// ...

	d.wrappee.ReadFile(inputFile)
	return d
}

// Populate Decrypt fields with decrypted text and error if it happened.
// Decrypts ciphertext from Reader.GetContent().
// Skips action if reader already has an error.
// func (d Decrypt) DoGenericReaderDecrypt(r Reader) {
// 	if r.GetError() != nil {
// 		return
// 	}

// 	text, err := d.DecryptFileKey(r.GetContent(), d.keyPath)

// 	d.resultText = text
// 	r.SetContent(text)
// 	d.resultErr = err
// 	r.SetError(err)
// }

// Setter for keyPath attribute
func (d Decrypt) SetKeyPath(keyPath string) Decrypt {
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
