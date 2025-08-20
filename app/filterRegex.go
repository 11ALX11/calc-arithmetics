package app

import (
	"fmt"
	"regexp"
)

/*
ReplaceMathExpressionsRegex searches input string for arithmetic exprs,
then replaces each one with result of an evaluation func. Uses regexp lib.
*/
func ReplaceMathExpressionsRegex(input string, evalFunc func(string) float64) string {

	// ToDo: make it recursive and use magic like (?=), which regexp lib doesn't support
	// https://github.com/google/re2/wiki/Syntax
	pattern := regexp.MustCompile(`\(*\d+(\s*[\+\-\*\/]\s*[\(\)]*\s*\d+[\)]*)+`)
	modifiedContent := pattern.ReplaceAllStringFunc(input, func(match string) string {
		return fmt.Sprintf("%.4g", evalFunc(match))
	})

	return modifiedContent
}
