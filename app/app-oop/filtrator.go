package app_oop

// Filtrator represents a type that can replace arithmetic expressions in a string.
type Filtrator interface {
	// Same as ReplaceMathExpressions() or ReplaceMathExpressionsRegex() in app package.
	// Attribute evalFunc needs to be set before calling this method
	ReplaceMathExpressions(input string) string

	// Setter for evalFunc attribute
	SetEvalFunc(evalFunc func(string) float64) Filtrator
	// Setter for evalFunc attribute
	SetEvalFuncWithEvaluator(evaluator Evaluator) Filtrator
}
