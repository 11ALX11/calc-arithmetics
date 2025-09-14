package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

type EvalLib struct{}

func (e EvalLib) Evaluate(expression string) float64 {
	return app.EvalLib(expression)
}
