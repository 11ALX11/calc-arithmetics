package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// EvalLib represents a type that can evaluate arithmetic expressions. Uses expr-lang lib.
type EvalLib struct{}

/*
Evaluate evaluates(calculates) an arithmetic expr. Uses expr-lang lib.

@param expr string - string with an arithmetic expr to evaluate.
String is allowed to contain only 0-9, +-/* and (), also expr itself needs to be correct.

@return float64 - result of an expr.
*/
func (e EvalLib) Evaluate(expression string) float64 {
	return app.EvalLib(expression)
}
