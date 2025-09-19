package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// EvalLib represents a type that can evaluate arithmetic expressions. Uses expr-lang lib.
type EvalLib struct{}

/*
Same as EvalLib() in app package
*/
func (e EvalLib) Evaluate(expression string) float64 {
	return app.EvalLib(expression)
}
