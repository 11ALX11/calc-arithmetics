package app

import (
	"fmt"
	"strings"

	"github.com/GRbit/go-pcre"
)

/*
ReplaceMathExpressionsRegex searches input string for arithmetic exprs,
then replaces each one with result of an evaluation func. Uses regexp lib.
*/
func ReplaceMathExpressionsRegex(input string, evalFunc func(string) int) string {
	// Define the PCRE expression (with extended syntax for better readability)
	pattern := `(?x)
		(?&expr)

		(?(DEFINE)
			(?<expr>
				[+-]*\(\s*(?&inner)\s*\)
				|
				[+-]*\s*(?&term)(?:\s*[+\-*\/]\s*(?&term))+
				|
				[+-]+\d+
			)
			(?<inner>
				(?&term)(?:\s*[+\-*\/]\s*(?&term))*
			)
			(?<term>
				(?&factor)(?:\s*[*\/]\s*(?&factor))*
			)
			(?<factor>
				\(\s*(?&inner)\s*\)
				| \d+
				| [+-]+\d+
				| [+-]+\(\s*(?&inner)\s*\)
			)
		)`

	// Compile the regular expression using pcre
	re := pcre.MustCompileJIT(pattern, pcre.EXTENDED, pcre.CONFIG_JIT)

	matcher := re.NewMatcherString(input, 0)
	matcher.ExecString(input, 0)

	matches := matcher.ExtractString() // error

	result := input

	for _, substr := range matches {
		newstr := fmt.Sprint(evalFunc(substr))
		result = strings.Replace(result, substr, newstr, 1)
	}

	return result
}
