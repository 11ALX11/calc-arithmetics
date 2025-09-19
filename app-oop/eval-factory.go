package app_oop

// EvalFactory is a factory that gives an Evaluator implementation based on the provided flag.
type EvalFactory struct {
	useEvalLib bool
}

// NewEvalFactory is a constructor for a factory that gives an Evaluator implementation based on the provided flag.
func NewEvalFactory(useEvalLib bool) *EvalFactory {
	e := new(EvalFactory)
	e.useEvalLib = useEvalLib
	return e
}

// GetEvalImplementation() is a factory method that returns an Evaluator implementation based on the provided flag.
func (e EvalFactory) GetEvalImplementation() Evaluator {
	if e.useEvalLib {
		return &EvalLib{}
	}
	return &Eval{}
}
