package app_oop

// Evaluator represents a type that can evaluate arithmetic expressions.
type Evaluator interface {
	// Same as Eval() or EvalLib() in app package
	Evaluate(expression string) float64
}
