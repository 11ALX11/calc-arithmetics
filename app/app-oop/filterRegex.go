package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Eval represents a type that can replace arithmetic expressions in a string..
type FilterRegex struct {
	evalFunc func(string) float64
}

// Setter for evalFunc attribute
func (f FilterRegex) SetEvalFunc(evalFunc func(string) float64) Filtrator {
	f.evalFunc = evalFunc
	return f
}

/*
Same as ReplaceMathExpressionsRegex() in app package. Attribute evalFunc needs to be set before calling this method
*/
func (f FilterRegex) ReplaceMathExpressions(input string) string {
	return app.ReplaceMathExpressionsRegex(input, f.evalFunc)
}
