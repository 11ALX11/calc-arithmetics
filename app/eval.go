package app

import (
	"log"
)

const (
	// EvalWrongOperand is returned by Eval when unknown operand is parsed.
	EvalWrongOperand = -101
	// EvalInvalidExpression is returned by Eval when len(numbers) at the end is isnt a 1
	EvalInvalidExpression = -102
	// EvalMismatchedParentheses is returned by Eval when parentheses are mismatched
	EvalMismatchedParentheses = -103
	// EvalInvalidUnaryOperator is returned by Eval when we get / or * as unaries
	EvalInvalidUnaryOperator = -104
	// EvalUnexpectedCharacter is returned by Eval when we encounter unexpected character during parsing
	EvalUnexpectedCharacter = -105
)

/*
Eval evaluates(calculates) an arithmetic expression.

@param expression string - string with an arithmetic expression to evaluate.
String is allowed to contain only 0-9, +-/* and (), also expression itself needs to be correct.

@return float64 - result of an expr.
*/
func Eval(expression string) float64 {
	precedence := func(operand byte) int {
		switch operand {
		case '(', ')':
			return 0
		case '+', '-':
			return 1
		default: // *, /
			return 2
		}
	}

	applyOperand := func(operand byte, b, a float64) float64 {
		switch operand {
		case '+':
			return a + b
		case '-':
			return a - b
		case '*':
			return a * b
		case '/':
			return a / b
		default:
			log.Println("Eval(): Wrong operand")
			return EvalWrongOperand
		}
	}

	var numbers []float64
	var operands []byte

	// helper to apply top operator once (assumes stacks are valid)
	applyTop := func() {
		number2 := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		number1 := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		operand := operands[len(operands)-1]
		operands = operands[:len(operands)-1]
		numbers = append(numbers, applyOperand(operand, number2, number1))
	}

	// isUnary context: true when the next + or - should be parsed as unary.
	// This is true at expression start, after '(', or after another operator.
	isUnary := true

	i := 0
	for i < len(expression) {
		ch := expression[i]

		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			i++
			continue
		}

		if ch >= '0' && ch <= '9' {
			// parse integer number
			number := 0.0
			for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
				number = number*10 + float64(expression[i]-'0')
				i++
			}
			numbers = append(numbers, number)
			isUnary = false
			continue // note: already advanced i
		}

		if ch == '(' {
			operands = append(operands, ch)
			isUnary = true
			i++
			continue
		}

		if ch == ')' {
			// pop until '('
			for len(operands) > 0 && operands[len(operands)-1] != '(' {
				applyTop()
			}
			if len(operands) == 0 {
				log.Println("Eval(): Mismatched parentheses.")
				return EvalMismatchedParentheses
			}
			operands = operands[:len(operands)-1] // pop '('
			isUnary = false
			i++
			continue
		}

		// Operators: + - * /
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			if isUnary {
				// handle unary +/-
				if ch == '+' {
					// unary plus: no-op, just keep expecting a number/parenthesis next
					// But to handle sequences like + - 1, we simply keep isUnary = true.
					i++
					// still unary until we actually read a number or '('
					continue
				}
				if ch == '-' {
					// unary minus: turn into "0 - (...)"
					// Push a zero, then treat '-' as binary with high precedence resolution below.
					numbers = append(numbers, 0.0)
					// Now treat '-' as binary subtraction: fall through to normal binary handling
					// but we must mark isUnary=false because we are placing an operator now expecting right operand.
					// We won't increment i here; let the normal operator handling below consume '-'.
					isUnary = true // remains true because after a binary operator we still expect an operand
				} else {
					// unary * or / is invalid in standard arithmetic
					log.Println("Eval(): Invalid unary operator.")
					return EvalInvalidUnaryOperator
				}
			}

			// Now handle binary operators (including the '-' that came from unary minus transformation)
			for len(operands) > 0 && operands[len(operands)-1] != '(' &&
				precedence(operands[len(operands)-1]) >= precedence(ch) {
				applyTop()
			}
			operands = append(operands, ch)
			isUnary = true // after a binary operator, next + or - could be unary
			i++
			continue
		}

		log.Printf("Eval(): Unexpected character: %q at %d\n", ch, i)
		return EvalUnexpectedCharacter
	}

	// apply remaining operators
	for len(operands) > 0 {
		if operands[len(operands)-1] == '(' {
			log.Println("Eval(): Mismatched parentheses.")
			return EvalMismatchedParentheses
		}
		applyTop()
	}

	if len(numbers) != 1 {
		log.Println("Eval(): Invalid expression.")
		return EvalInvalidExpression
	}

	return numbers[0]
}
