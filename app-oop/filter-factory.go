package app_oop

// FilterFactory is a factory that gives an Filtrator implementation based on the provided flag.
type FilterFactory struct {
	useFilterRegex bool
}

// NewFilterFactory is a constructor for a factory that gives an Filtrator implementation based on the provided flag.
func NewFilterFactory(useFilterRegex bool) *FilterFactory {
	f := new(FilterFactory)
	f.useFilterRegex = useFilterRegex
	return f
}

// GetFilterImplementation() is a factory method that returns an Filtrator implementation based on the provided flag.
func (f FilterFactory) GetFilterImplementation(evaluator Evaluator) Filtrator {
	if f.useFilterRegex {
		return (&FilterRegex{}).SetEvalFuncWithEvaluator(evaluator)
	}
	return (&Filter{}).SetEvalFuncWithEvaluator(evaluator)
}
