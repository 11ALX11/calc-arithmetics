package app

import (
	"fmt"
	"strings"

	"github.com/GRbit/go-pcre"
)

// Define the PCRE expression
const pcrePattern = `(?mx)
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
var regexpPcre pcre.Regexp = pcre.MustCompileJIT(pcrePattern, pcre.MULTILINE|pcre.EXTENDED, pcre.CONFIG_JIT)
var matcherPcre pcre.Matcher = *regexpPcre.NewMatcherString("", pcre.MULTILINE|pcre.EXTENDED)

/*
ReplaceMathExpressionsRegex searches input string for arithmetic exprs,
then replaces each one with result of an evaluation func. Uses regexp lib.
*/
func ReplaceMathExpressionsRegex(input string, evalFunc func(string) int) string {

	matcherPcre.ExecString(input, pcre.MULTILINE|pcre.EXTENDED)
	matches := matcherPcre.ExtractString()

	// No syntax errors, but nothing is matching

	result := input

	for _, substr := range matches {
		newstr := fmt.Sprint(evalFunc(substr))
		result = strings.Replace(result, substr, newstr, 1)
	}

	return result
}
