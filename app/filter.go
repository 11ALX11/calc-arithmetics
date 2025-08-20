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
func ReplaceMathExpressions(input string, evalFunc func(string) float64) string {

	result := input

	for _, substr := range extractMathExpressions(input) {
		newstr := fmt.Sprintf("%.4g", evalFunc(substr))
		result = strings.Replace(result, substr, newstr, 1)
	}

	return result
}

// extractMathExpressions extracts arithmetic expressions from a given input string.
func extractMathExpressions(s string) []string {
	var result []string

	var expr strings.Builder // Buffer for an expression
	inExpr := false          // Flag indicating when we iterate through an expression
	parentheses := 0         // Parentheses pairs state counter
	startExprInd := 0        // Index where current expression started
	lastMathRune := 'd'      // Last mathematical character seen ('d' default, before we passed first character)
	lastDigitOperRune := 'd' // Last digit or operator seen
	possibleEnd := false     // Flag indicating possible expression end
	faultOperInd := -1       // Index of faulty operator
	faultyOperator := false  // Flag for invalid operator placement

	// Append spaces to simplify string handling
	s = " " + s + " "

	i := 0
	for i < len(s) {
		r := rune(s[i])

		if (unicode.IsDigit(r) || strings.ContainsRune("+-*/()", r)) && i != faultOperInd {
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
				if strings.ContainsRune("d(+-*/", lastMathRune) {
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

			if i == faultOperInd {
				faultOperInd = -1
			}

			trimmed := strings.TrimSpace(expr.String())

			// Check if operator doesnt have digits after it at the end of an expression.
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
			} else {
				// end of an expression

				// Check if theres '(' somewhere in expression
				if len(trimmed) > 0 && strings.ContainsAny(trimmed, "+-*/") && parentheses > 0 {
					faultyParenthesisInd := startExprInd + strings.IndexRune(
						s[startExprInd:i],
						'(',
					)

					// Check if this '(' isn't in a begining of an expression
					if faultyParenthesisInd > startExprInd {
						// Divide expression.
						// P.S. Shouldn't cause problem with operators, since we already checked them
						// and should there be any, we are either triggered by them first
						// or parenthesis comes first
						faultOperInd = faultyParenthesisInd
					}
				}

				if faultOperInd > startExprInd {
					i = startExprInd - 1
				} else {
					i = startExprInd
				}
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
