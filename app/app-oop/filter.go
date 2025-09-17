package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Filter represents a type that can replace arithmetic expressions in a string.
type Filter struct {
	evalFunc func(string) float64
}

// Setter for evalFunc attribute
func (f Filter) SetEvalFunc(evalFunc func(string) float64) Filtrator {
	f.evalFunc = evalFunc
	return f
}

// Setter for evalFunc attribute
func (f Filter) SetEvalFuncWithEvaluator(evaluator Evaluator) Filtrator {
	f.evalFunc = evaluator.Evaluate
	return f
}

/*
Same as ReplaceMathExpressions() in app package. Attribute evalFunc needs to be set before calling this method
*/
func (f Filter) ReplaceMathExpressions(input string) string {
	return app.ReplaceMathExpressions(input, f.evalFunc)
}
