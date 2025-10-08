package app_oop

// FilterFactory is a factory that gives an Filtrator implementation based on the provided flag.
type FilterFactory struct {
	useFilterRegex bool
}

// NewFilterFactory is a constructor for a factory that gives an Filtrator implementation based on the provided flag.
func NewFilterFactory(useFilterRegex bool) *FilterFactory {
	return &FilterFactory{useFilterRegex}
}

// GetFilterImplementation() is a factory method that returns an Filtrator implementation based on the provided flag.
func (f FilterFactory) GetFilterImplementation(evaluator Evaluator) Filtrator {

	var filtrator Filtrator = &Filter{}
	if f.useFilterRegex {
		filtrator = &FilterRegex{}
	}

	// return filtrator.SetEvalFuncWithEvaluator(evaluator) doesn't work for some reason ¯\_(ツ)_/¯.
	// So this is a workaround, since setter doesn't create new Filtrator,
	// but modifies existing one (and returns existing one) (P.S. as it should).
	filtrator.SetEvalFuncWithEvaluator(evaluator)
	return filtrator
}
