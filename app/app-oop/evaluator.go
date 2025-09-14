package app_oop

// Evaluator represents a type that can evaluate arithmetic expressions.
type Evaluator interface {
	Evaluate(expression string) float64
}
