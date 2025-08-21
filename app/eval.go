package app

/*
Eval evaluates(calculates) an arithmetic expression.

@param expression string - string with an arithmetic expression to evaluate.
String is allowed to contain only 0-9, +-/* and (), also expression itself needs to be correct.

@return int - result of an expr.
*/
func Eval(expression string) float64 {
	var precedence = func(operand byte) int {
		switch operand {
		case '(', ')':
			return 0
		case '+', '-':
			return 1
		}
		return 2
	}

	var applyOperand = func(operand byte, b, a float64) float64 {
		switch operand {
		case '+':
			return a + b
		case '-':
			return a - b
		case '*':
			return a * b
		case '/':
			return a / b
		}
		return 0
	}

	var numbers []float64
	var operands []byte

	i := 0
	for i < len(expression) {
		if expression[i] == ' ' {
			i++
			continue
		} else if expression[i] >= '0' && expression[i] <= '9' {
			number := 0.
			for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
				number = number*10 + float64(expression[i]-'0')
				i++
			}
			numbers = append(numbers, number)
			i--
		} else if expression[i] == '(' {
			operands = append(operands, expression[i])
		} else if expression[i] == ')' {
			for operands[len(operands)-1] != '(' {
				number2 := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				number1 := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				operand := operands[len(operands)-1]
				operands = operands[:len(operands)-1]
				numbers = append(numbers, applyOperand(operand, number2, number1))
			}
			operands = operands[:len(operands)-1]
		} else {
			for len(operands) > 0 && operands[len(operands)-1] != '(' && precedence(operands[len(operands)-1]) >= precedence(expression[i]) {
				number2 := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				number1 := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				operand := operands[len(operands)-1]
				operands = operands[:len(operands)-1]
				numbers = append(numbers, applyOperand(operand, number2, number1))
			}
			operands = append(operands, expression[i])
		}
		i++
	}

	for len(operands) > 0 {
		number2 := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		number1 := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]
		operand := operands[len(operands)-1]
		operands = operands[:len(operands)-1]
		numbers = append(numbers, applyOperand(operand, number2, number1))
	}

	return numbers[0]
}
