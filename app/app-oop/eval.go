package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Eval represents a type that can evaluate arithmetic expressions.
type Eval struct{}

/*
Evaluate evaluates (calculates) an arithmetic expression.

Allowed tokens: digits 0-9, + - * / and parentheses ().
Supports unary + and - (e.g., -1, --1, + - 1, -(-1), etc.).

@return float64 - result of the expression, or a negative error code on parse/semantic error.
*/
func (e Eval) Evaluate(expression string) float64 {
	return app.Eval(expression)
}
