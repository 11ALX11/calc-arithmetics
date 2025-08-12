package app

import (
	"log"
	"math"

	"github.com/expr-lang/expr"
)

const (
	// EvalLibEvaluationError is returned by EvalLib when expr-lang fails to compile the expression.
	EvalLibCompilationError = -11
	// EvalLibRunError is returned by EvalLib when expr-lang fails to run the bytecode
	EvalLibRunError = -12
)

/*
EvalLib evaluates(calculates) an arithmetic expr. Uses expr-lang lib.

@param expr string - string with an arithmetic expr to evaluate.
String is allowed to contain only 0-9, +-/* and (), also expr itself needs to be correct.

@return int - result of an expr, rounded to the nearest integer.
*/
func EvalLib(expression string) int {

	bytecode, err1 := expr.Compile(expression, expr.AsFloat64())
	if err1 != nil {
		log.Printf("failed to compile an expression %s: %v", expression, err1)
		return EvalLibCompilationError
	}

	out, err2 := expr.Run(bytecode, nil)
	if err2 != nil {
		log.Printf("failed to run an expression %s: %v", expression, err2)
		return EvalLibRunError
	}

	result := int(math.Round(out.(float64)))

	return result
}
