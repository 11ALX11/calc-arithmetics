package app_oop

// NewEvaluator is a factory function that returns an Evaluator implementation based on the provided flag.
func NewEvaluator(useEvalLib bool) Evaluator {
	if useEvalLib {
		return EvalLib{}
	}
	return Eval{}
}
