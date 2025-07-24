package app

/*
Eval evaluates(calculates) an arithmetic expr.

@param expr string - string with an arithmetic expr to evaluate.
String is allowed to contain only 0-9, +-/* and (), also expr itself needs to be correct.

@return int - result of an expr.
*/
func Eval(expr string) int {
	var precedence = func(operand byte) int {
		if operand == '(' || operand == ')' {
			return 0
		} else if operand == '+' || operand == '-' {
			return 1
		}
		return 2
	}

	var applyOperand = func(operand byte, b, a int) int {
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

	var numbers []int
	var operands []byte

	i := 0
	for i < len(expr) {
		if expr[i] == ' ' {
			i++
			continue
		} else if expr[i] >= '0' && expr[i] <= '9' {
			number := 0
			for i < len(expr) && expr[i] >= '0' && expr[i] <= '9' {
				number = number*10 + int(expr[i]-'0')
				i++
			}
			numbers = append(numbers, number)
			i--
		} else if expr[i] == '(' {
			operands = append(operands, expr[i])
		} else if expr[i] == ')' {
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
			for len(operands) > 0 && operands[len(operands)-1] != '(' && precedence(operands[len(operands)-1]) >= precedence(expr[i]) {
				number2 := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				number1 := numbers[len(numbers)-1]
				numbers = numbers[:len(numbers)-1]
				operand := operands[len(operands)-1]
				operands = operands[:len(operands)-1]
				numbers = append(numbers, applyOperand(operand, number2, number1))
			}
			operands = append(operands, expr[i])
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
