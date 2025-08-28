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
Eval evaluates (calculates) an arithmetic expression.

Allowed tokens: digits 0-9, + - * / and parentheses ().
Supports unary + and - (e.g., -1, --1, + - 1, -(-1), etc.).

@return float64 - result of the expression, or a negative error code on parse/semantic error.
*/
func Eval(expression string) float64 {
	// Precedence: higher number = higher precedence
	// Assign unary operators higher precedence than * and /
	precedence := func(operand byte) int {
		switch operand {
		case '(', ')':
			return 0
		case 'p', 'n': // unary plus, unary minus
			return 3
		case '*', '/':
			return 2
		case '+', '-':
			return 1
		default:
			return -1
		}
	}

	// Associativity: true if right-associative
	isRightAssoc := func(operand byte) bool {
		switch operand {
		case 'p', 'n':
			return true // unary operators are right-associative
		default:
			return false
		}
	}

	applyBinary := func(operand byte, b, a float64) float64 {
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

	applyUnary := func(operand byte, a float64) float64 {
		switch operand {
		case 'p': // unary plus
			return a
		case 'n': // unary minus
			return -a
		default:
			log.Println("Eval(): Wrong unary operand")
			return EvalWrongOperand
		}
	}

	var numbers []float64
	var operands []byte

	applyTop := func() {
		op := operands[len(operands)-1]
		operands = operands[:len(operands)-1]

		if op == 'p' || op == 'n' {
			// unary
			if len(numbers) < 1 {
				log.Println("Eval(): Invalid expression (unary needs 1 operand)")
				numbers = append(numbers, EvalInvalidExpression)
				return
			}
			a := numbers[len(numbers)-1]
			numbers = numbers[:len(numbers)-1]
			numbers = append(numbers, applyUnary(op, a))
			return
		}

		// binary
		if len(numbers) < 2 {
			log.Println("Eval(): Invalid expression (binary needs 2 operands)")
			numbers = append(numbers, EvalInvalidExpression)
			return
		}
		b := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		a := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		numbers = append(numbers, applyBinary(op, b, a))
	}

	// isUnary context: true when the next + or - should be parsed as unary.
	isUnary := true

	i := 0
	for i < len(expression) {
		ch := expression[i]

		// skip whitespace
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			i++
			continue
		}

		// number
		if ch >= '0' && ch <= '9' {
			number := 0.0
			for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
				number = number*10 + float64(expression[i]-'0')
				i++
			}
			numbers = append(numbers, number)
			isUnary = false
			continue
		}

		// left paren
		if ch == '(' {
			operands = append(operands, ch)
			isUnary = true
			i++
			continue
		}

		// right paren
		if ch == ')' {
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

		// operators
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			// classify unary vs binary
			op := ch
			if isUnary {
				// only + and - are valid unary
				switch ch {
				case '+':
					op = 'p' // unary plus
				case '-':
					op = 'n' // unary minus
				default:
					log.Println("Eval(): Invalid unary operator.")
					return EvalInvalidUnaryOperator
				}
			}

			// pop while top of operator stack has greater precedence,
			// or equal precedence and current operator is left-associative.
			for len(operands) > 0 && operands[len(operands)-1] != '(' {
				top := operands[len(operands)-1]
				pt := precedence(top)
				po := precedence(op)
				if pt > po || (pt == po && !isRightAssoc(op)) {
					applyTop()
				} else {
					break
				}
			}

			operands = append(operands, op)
			// after any operator, we are in unary context again
			isUnary = true
			i++
			continue
		}

		log.Printf("Eval(): Unexpected character: %q at %d\n", ch, i)
		//return EvalUnexpectedCharacter // uh, just ignore it
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
