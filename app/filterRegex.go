package app

import (
	"fmt"
	"log"
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
var regexpPcre pcre.Regexp = pcre.MustCompileJIT(pcrePattern, pcrePatternFlags, pcre.CONFIG_JIT)
var matcherPcre pcre.Matcher = *regexpPcre.NewMatcherString("", pcrePatternFlags)

/*
ReplaceMathExpressionsRegex searches input string for arithmetic exprs,
then replaces each one with result of an evaluation func. Uses regexp lib.
*/
func ReplaceMathExpressionsRegex(input string, evalFunc func(string) int) string {

	// matcherPcre.ExecString(input, pcrePatternFlags)
	// matches := matcherPcre.ExtractString()

	// No syntax errors, but nothing is matching

	result := input
	offset := 0

	for offsetFromExec := matcherPcre.ExecString(input[offset:], pcre.STUDY_JIT_COMPILE); offsetFromExec > 0; {

		matches := matcherPcre.ExtractString()
		var match string

		// log.Printf("Matcher: %v", matcherPcre)

		if len(matches) > 0 {
			match = matches[0]
			log.Printf("Match: %v", match)
		} else {
			continue
		}

		newstr := fmt.Sprint(evalFunc(match))
		result = strings.Replace(result, match, newstr, 1)

		offset += matcherPcre.Index()[1] // advance past this match
	}

	return result
}
