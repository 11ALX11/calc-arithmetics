package app_oop

// Evaluator represents any object with an ability to evaluate arithmetic expression.
type Evaluator interface {
	Evaluate(expression string) float64
}
