package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Filter represents a type that can replace arithmetic expressions in a string.
type Filter struct {
	IFiltrator
}

/*
Same as ReplaceMathExpressions() in app package.
Assumes evalFunc has been set via SetEvalFunc or SetEvalFuncWithEvaluator (as of 14.10.2025: guaranteed by factory).
*/
func (f Filter) ReplaceMathExpressions(input string) string {
	return app.ReplaceMathExpressions(input, f.evalFunc)
}
