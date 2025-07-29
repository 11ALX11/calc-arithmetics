package app

import (
	"fmt"
	"log"
	"strconv"

	"github.com/expr-lang/expr"
)

/*
EvalLib evaluates(calculates) an arithmetic expr. Uses expr-lang lib.

@param expr string - string with an arithmetic expr to evaluate.
String is allowed to contain only 0-9, +-/* and (), also expr itself needs to be correct.

@return int - result of an expr.
*/
func EvalLib(expression string) int {
	out, err := expr.Eval(expression, nil)
	if err != nil {
		log.Fatalf("failed to evaluate an expression %s: %v", expression, err)
	}

	result, err := strconv.Atoi(fmt.Sprint(out))
	if err != nil {
		log.Printf("failed to convert result of an evaluation of %s to int: %v", expression, err)
		return -1
	}

	return result
}
