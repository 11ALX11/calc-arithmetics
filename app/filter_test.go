package app

import (
	"strings"
	"testing"
)

var filterTests = []struct {
	in  string
	out string
}{
	{"Here's an arithmetic 1+1=2.", "Here's an arithmetic 2=2."},
	{"Here's another one: (2+6* 3+5- (3*14/7+2)*5)+3 = -12.", "Here's another one: -12 = -12."},
}

func TestFilter(t *testing.T) {

	var tests []struct {
		in  string
		out string
	}

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			str := strings.Trim(ReplaceMathExpressions(tt.in), " ")
			if str != tt.out {
				t.Errorf("got %s, want %s", str, tt.out)
			}
		})
	}
}
