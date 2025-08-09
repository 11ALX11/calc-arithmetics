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

// isMathChar checks if a rune is a digit, arithmetic operator or parenthesis.
// Also checks if closing parenthesis ')' is valid
func isMathChar(r rune) bool {
	return unicode.IsDigit(r) || strings.ContainsRune("+-*/()", r)
}

// extractMathExpressions extracts arithmetic expressions from a given input string.
// It doesnt work at all...
func extractMathExpressions(s string) []string {
	var result []string
	var expr strings.Builder
	inExpr := false
	parentheses := 0
	startExprInd := 0
	lastMathRune := 'd'
	possibleEnd := false

	i := 0
	for i < len(s) {
		r := rune(s[i])

		if isMathChar(r) && parentheses >= 0 {
			if !inExpr {
				// Start of a new expression
				inExpr = true
				expr.Reset()
				parentheses = 0
				startExprInd = i
				lastMathRune = 'd'
			}

			if r == '(' {
				if unicode.IsDigit(r) || lastMathRune == ')' {
					possibleEnd = true
				}

				parentheses++
			} else if r == ')' {
				if strings.ContainsRune("(+-/*", lastMathRune) || parentheses == 0 {
					possibleEnd = true
				}

				parentheses--
			} else if strings.ContainsRune("+-*/", lastMathRune) {
				if lastMathRune == '(' {
					possibleEnd = true
				}
			} else if unicode.IsDigit(r) {
				if lastMathRune == ')' {
					possibleEnd = true
				}
			}

		} else {
			// Non-math character, possibly end of an expression
			possibleEnd = true
		}

		// Handle end of string
		if i == len(s)-1 && inExpr {
			possibleEnd = true
		}

		if possibleEnd && inExpr {

			trimmed := strings.TrimSpace(expr.String())
			if len(trimmed) > 0 && (parentheses == 0) && strings.ContainsAny(trimmed, "+-*/") {

				result = append(result, trimmed)

				// revisit and start new expr from current rune
				// only if its not ')'
				if r != ')' {
					i--
				}

			} else {
				// end of an expression,
				// but its 0 length or invalid parenthesis
				i = startExprInd
			}

			inExpr = false
		}
		possibleEnd = false

		if inExpr {
			if unicode.IsSpace(r) {
				// Allow spaces *within* expressions but not at the start
				// To prevent splitting expressions like "2 + 2"
				expr.WriteRune(' ')
			} else {
				expr.WriteRune(r)
			}
		}

		lastMathRune = r
		i++
	}

	return result
}
