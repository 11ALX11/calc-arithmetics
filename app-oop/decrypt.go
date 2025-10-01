package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Decrypt is a type that decrypts a ciphertext string using AES. Implements ReaderDecorator
type Decrypt struct {
	wrappee Reader

	keyPath string

	resultText string
	resultErr  error
}

// NewDecrypt is a constructor for a Decrypt.
func NewDecrypt(reader Reader) *Decrypt {
	f := new(Decrypt)
	f.wrappee = reader
	return f
}

// Setter for dataInputFile attribute
func (r *Decrypt) SetDataInputFile(dataInputFile string) Reader {
	return r.wrappee.SetDataInputFile(dataInputFile)
}

// Getter for content attribute
func (r Decrypt) GetContent() string {
	return r.wrappee.GetContent()
}

// Setter for content attribute
func (r *Decrypt) SetContent(s string) Reader {
	return r.wrappee.SetContent(s)
}

// Getter for err attribute
func (r Decrypt) GetError() error {
	return r.wrappee.GetError()
}

// Setter for content attribute
func (r *Decrypt) SetError(e error) Reader {
	return r.wrappee.SetError(e)
}

// Getter for both content and error.
// Ex: content, err := reader.GetContentError()
func (r Decrypt) GetContentError() (string, error) {
	return r.wrappee.GetContentError()
}

/*
Same as ReadFile() in app package
*/
func (r *Decrypt) ReadFile(inputFile string) Reader {

	// ...

	return r.wrappee.ReadFile(inputFile)
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
