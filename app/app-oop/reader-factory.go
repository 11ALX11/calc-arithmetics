package app_oop

// EvalFactory is a factory that gives an Evaluator implementation based on the provided flag.
type ReaderFactory struct {
	unzip bool
}

// NewEvalFactory is a constructor for a factory that gives an Evaluator implementation based on the provided flag.
func NewReaderFactory(unzip bool) *ReaderFactory {
	r := new(ReaderFactory)
	r.unzip = unzip
	return r
}

// GetEvalImplementation() is a factory method that returns an Evaluator implementation based on the provided flag.
func (r ReaderFactory) GetReaderImplementation() Reader {
	if r.unzip {
		return &ReadinUnzip{}
	}
	return &Readin{}
}
