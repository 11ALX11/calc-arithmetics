package app_oop

import (
	"github.com/11ALX11/calc-arithmetics/app"
)

// Decrypt is a type that decrypts a ciphertext string using AES. Implements ReaderVisitor
type Decrypt struct {
	keyPath string

	resultText string
	resultErr  error
}

func NewDecrypt() *Decrypt {
	d := new(Decrypt)
	return d
}

// Visit Readin
func (d Decrypt) DoForReadin(r *Readin) {
	d.DoGenericReaderDecrypt(r)
}

// Visit ReadinUnzip
func (d Decrypt) DoForReadinUnzip(r *ReadinUnzip) {
	d.DoGenericReaderDecrypt(r)
}

// Populate Decrypt fields with decrypted text and error if it happened.
// Decrypts ciphertext from Reader.GetContent().
// Skips action if reader already has an error.
func (d Decrypt) DoGenericReaderDecrypt(r Reader) {
	if r.GetError() != nil {
		return
	}

	text, err := d.DecryptFileKey(r.GetContent(), d.keyPath)

	d.resultText = text
	r.SetContent(text)
	d.resultErr = err
	r.SetError(err)
}

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
