package app_oop

import "github.com/11ALX11/calc-arithmetics/app"

// Eval represents a type that can replace arithmetic expressions in a string..
type FilterRegex struct {
	IFiltrator
}

/*
Same as ReplaceMathExpressionsRegex() in app package. Attribute evalFunc needs to be set before calling this method
*/
func (f FilterRegex) ReplaceMathExpressions(input string) string {
	return app.ReplaceMathExpressionsRegex(input, f.evalFunc)
}
