package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Eval represents a type that can evaluate arithmetic expressions.
type Eval struct{}

/*
Same as Eval() in app package
*/
func (e Eval) Evaluate(expression string) float64 {
	return app.Eval(expression)
}
