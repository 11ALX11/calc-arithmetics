package app

import (
	"fmt"
	"strings"

	"github.com/GRbit/go-pcre"
)

// PCRE flags "mx"
const pcrePatternFlags = pcre.MULTILINE | pcre.EXTENDED

// Define the PCRE expression
const pcrePattern = `(?mx)
		(?&expr)

		(?(DEFINE)
			(?<expr>
				[+-]*\s*(?&term)(?:\s*[+\-*\/]\s*(?&term))+
				|
				[+-]*\s*\(\s*(?&inner)\s*\)
				|
				(?:[+-]\s*)+\d+
			)
			(?<inner>
				(?&term)(?:\s*[+\-*\/]\s*(?&term))*
			)
			(?<term>
				(?&factor)(?:\s*[*\/]\s*(?&factor))*?
			)
			(?<factor>
				\(\s*(?&inner)\s*\)
				| \d+
				| [+-]+\s*\d+
				| [+-]+\(\s*(?&inner)\s*\)
			)
		)`

// Compile the regular expression using pcre
var regexpPcre pcre.Regexp = pcre.MustCompileJIT(pcrePattern, pcrePatternFlags, pcre.CONFIG_JIT)

/*
ReplaceMathExpressionsRegex searches input string for arithmetic exprs,
then replaces each one with result of an evaluation func. Uses regexp lib.
*/
func ReplaceMathExpressionsRegex(input string, evalFunc func(string) float64) string {

	result := input
	offset := 0

	matcher := regexpPcre.NewMatcherString(input, 0)
	for matcher.Matches {

		match := strings.TrimSpace(matcher.GroupString(0))

		newstr := fmt.Sprintf("%.4g", evalFunc(match))
		result = strings.Replace(result, match, newstr, 1)

		// Get next match (next iteration)
		offset += matcher.Index()[1]
		matcher = regexpPcre.NewMatcherString(input[offset:], 0)
	}

	return result
}
