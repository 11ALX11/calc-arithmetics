package app

import (
	"fmt"
	"strings"
	"unicode"
)

/*
ReplaceMathExpressions searches input string for arithmetic exprs,
then replaces each one with result of an evaluation func.
*/
func ReplaceMathExpressions(input string, evalFunc func(string) int) string {

	result := input

	for _, substr := range extractMathExpressions(input) {
		newstr := fmt.Sprint(evalFunc(substr))
		result = strings.Replace(result, substr, newstr, 1)
	}

	return result
}

// isMathChar checks if a rune is an arithmetic operator or parenthesis
func isMathChar(r rune) bool {
	return unicode.IsDigit(r) || strings.ContainsRune("+-*/()", r)
}

// extractMathExpressions extracts arithmetic expressions from a given input string.
//
// ToDo: manipulate i to return to a begining of an expr if ( didn't got closed.
// As of now, (1+1,2+4) breaks
// Also expr can end with ) if parentheses is 0
func extractMathExpressions(s string) []string {
	var result []string
	var expr strings.Builder
	inExpr := false
	parentheses := 0

	for i, r := range s {
		if isMathChar(r) {
			if !inExpr {
				// Start of a new expression
				inExpr = true
				expr.Reset()
				parentheses = 0
			}
			expr.WriteRune(r)
			switch r {
			case '(':
				parentheses++
			case ')':
				parentheses--
			}
		} else if unicode.IsSpace(r) && inExpr {
			// Allow spaces *within* expressions but not at the start
			// To prevent splitting expressions like "2 + 2"
			expr.WriteRune(' ')
		} else {
			// Non-math character, possibly end of an expression
			if inExpr {
				trimmed := strings.TrimSpace(expr.String())
				if len(trimmed) > 0 && (parentheses == 0) {
					// Avoid single numbers
					if strings.ContainsAny(trimmed, "+-*/") {
						result = append(result, trimmed)
					}
				}
				inExpr = false
			}
		}

		// Handle end of string
		if i == len(s)-1 && inExpr {
			trimmed := strings.TrimSpace(expr.String())
			if len(trimmed) > 0 && (parentheses == 0) {
				if strings.ContainsAny(trimmed, "+-*/") {
					result = append(result, trimmed)
				}
			}
			inExpr = false
		}
	}

	return result
}
