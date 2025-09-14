package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

type Eval struct{}

func (e Eval) Evaluate(expression string) float64 {
	return app.Eval(expression)
}
