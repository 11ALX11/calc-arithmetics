package app

import (
	"fmt"
	"regexp"
)

/*
ReplaceMathExpressions searches input string for arithmetic exprs,
then replaces each one with result of an evaluation func.
*/
func ReplaceMathExpressions(input string) string {

	pattern := regexp.MustCompile(`\(*\d+(\s*[\+\-\*\/]\s*[\(\)]*\s*\d+[\)]*)+`)
	modifiedContent := pattern.ReplaceAllStringFunc(input, func(match string) string {
		return fmt.Sprintf("%d", Eval(match))
	})

	return modifiedContent
}
