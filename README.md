# Visitor pattern example

This branch has been created to preserve an example of the visitor pattern in golang due to impending change.
I have plans to change Decrypt and other classes to implement decorator pattern for i/o operations.

Pattern usage: [cmd/cli.go:147](https://github.com/11ALX11/calc-arithmetics/blob/667b11841ef15f6870a43430026e7a3c1f4be0aa/cmd/cli.go#L147):

```
reader.Accept(app_oop.
	NewDecrypt().
	SetKeyPath(keyPath),
)

content, err := reader.GetContentError()
```

Visitor interface: [app_oop/reader-visitor:4](https://github.com/11ALX11/calc-arithmetics/blob/667b11841ef15f6870a43430026e7a3c1f4be0aa/app-oop/reader-visitor.go#L4):

```
// ReaderVisitor represents a type that implements visitor pattern to visit Reader interface
type ReaderVisitor interface {
	DoForReadin(r *Readin)
	DoForReadinUnzip(r *ReadinUnzip)
}
```

Reader (Component that accepts) interface: [app_oop/reader:4](https://github.com/11ALX11/calc-arithmetics/blob/667b11841ef15f6870a43430026e7a3c1f4be0aa/app-oop/reader.go#L4):

```
// Reader represents a type that can read from a file.
type Reader interface {

	// Accepts ReaderVisitor implementations and calls doFor*Implementation*()
	Accept(v ReaderVisitor)

	ReadFile(inputFile string) Reader

	SetDataInputFile(dataInputFile string) Reader
	GetContent() string
	SetContent(content string) Reader
	GetError() error
	SetError(err error) Reader
	GetContentError() (string, error)
}
```

Visitor implementation: [app_oop/decrypt:8](https://github.com/11ALX11/calc-arithmetics/blob/667b11841ef15f6870a43430026e7a3c1f4be0aa/app-oop/decrypt.go#L8):

```
// Decrypt is a type that decrypts a ciphertext string using AES. Implements ReaderVisitor
type Decrypt struct {
	keyPath string

	resultText string
	resultErr  error
}

// Constructor for Decrypt
func NewDecrypt() *Decrypt {
	d := new(Decrypt)
	return d
}

func (d Decrypt) DoForReadin(r *Readin) {
	d.DoGenericReaderDecrypt(r)
}

func (d Decrypt) DoForReadinUnzip(r *ReadinUnzip) {
	d.DoGenericReaderDecrypt(r)
}

// Populate Decrypt and Reader fields with decrypted text and error if it happened.
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
//...
```

Reader implementation 1: Readin: [app_oop/readin.go:8](https://github.com/11ALX11/calc-arithmetics/blob/667b11841ef15f6870a43430026e7a3c1f4be0aa/app-oop/readin.go#L8):

```
// Readin represents a type that can read a text file.
type Readin struct {
	content string
	err     error
}

// Accepts ReaderVisitor implementations and calls DoForReadin()
func (r *Readin) Accept(v ReaderVisitor) {
	v.DoForReadin(r)
}
//...
```

Reader implementation 2: ReadinUnzip: [app_oop/readinUnzip.go:6](https://github.com/11ALX11/calc-arithmetics/blob/667b11841ef15f6870a43430026e7a3c1f4be0aa/app-oop/readinUnzip.go#L6):

```
// ReadinUnzip represents a type that can read a zip file.
type ReadinUnzip struct {
	dataInputFile string

	content string
	err     error
}

// Accepts ReaderVisitor implementations and calls DoForReadinUnzip()
func (r *ReadinUnzip) Accept(v ReaderVisitor) {
	v.DoForReadinUnzip(r)
}
//...
```