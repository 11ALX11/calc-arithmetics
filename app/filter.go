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

// extractMathExpressions extracts arithmetic expressions from a given input string.
func extractMathExpressions(s string) []string {
	var result []string
	var expr strings.Builder
	inExpr := false
	parentheses := 0
	startExprInd := 0
	lastMathRune := 'd'
	lastDigitOperRune := 'd'
	possibleEnd := false
	faultOperInd := -1
	faultyOperator := false

	// simple fix to a check at an end of a line.
	s = s + " "

	i := 0
	for i < len(s) {
		r := rune(s[i])

		if unicode.IsDigit(r) || strings.ContainsRune("+-*/()", r) {
			if !inExpr {
				// Start of a new expression
				inExpr = true
				expr.Reset()
				parentheses = 0
				startExprInd = i
				lastMathRune = 'd'
				lastDigitOperRune = 'd'
			}

			if r == '(' {
				if unicode.IsDigit(lastMathRune) || lastMathRune == ')' {
					possibleEnd = true
				} else {
					parentheses++
				}
			} else if r == ')' {
				if strings.ContainsRune("(+-/*", lastMathRune) || parentheses == 0 {
					possibleEnd = true
				} else {
					parentheses--
				}
			} else if strings.ContainsRune("*/", r) {
				if lastMathRune == '(' {
					possibleEnd = true
				}
				lastDigitOperRune = r
			} else if strings.ContainsRune("+-", r) {
				lastDigitOperRune = r
			} else if unicode.IsDigit(r) {
				if lastMathRune == ')' {
					possibleEnd = true
				}
				lastDigitOperRune = r
			}

			lastMathRune = r
		} else if !unicode.IsSpace(r) {
			// Non-math character, except for spaces
			// Possibly end of an expression
			possibleEnd = true
		}

		// Handle end of string
		if i == len(s)-1 && inExpr {
			possibleEnd = true
		}

		if (possibleEnd || i == faultOperInd) && inExpr {

			trimmed := strings.TrimSpace(expr.String())

			if strings.ContainsRune("+-*/", lastDigitOperRune) {
				faultOperInd = strings.LastIndex(
					s[0:i],
					string(lastDigitOperRune),
				)

				faultyOperator = true
			}

			if len(trimmed) > 0 && parentheses == 0 && strings.ContainsAny(trimmed, "+-*/") && !faultyOperator {

				result = append(result, trimmed)

				// revisit and start new expr from current rune
				// only if its not ')'
				if r != ')' {
					i--
				}

				if i == faultOperInd {
					faultOperInd = -1
				}

			} else {
				// end of an expression
				i = startExprInd
			}

			inExpr = false
			parentheses = 0
			faultyOperator = false
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

		i++
	}

	return result
}
