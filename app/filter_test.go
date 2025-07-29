package app

import (
	"strings"
	"testing"
)

// ToDo: Add more tests to destroy what is built, attention to edge cases
var filterTests = []struct {
	in  string
	out string
}{
	{"Here's an arithmetic 1+1=2.", "Here's an arithmetic 2=2."},
	{"Here's another one: (2+6* 3+5- (3*14/7+2)*5)+3 = -12.", "Here's another one: -12 = -12."},
	{"Here's multiples: (2+6* 3+5- (3*14/7+2)*5)+3 = -12, 1 + 2 + 3 + 4 - 5 = 5", "Here's multiples: -12 = -12, 5 = 5"},
	{"Here's an arithmetic (1+1,2+4) 1+1=2.", "Here's an arithmetic (2,6) 2=2."},
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
			t.Parallel()
			str := strings.Trim(ReplaceMathExpressions(tt.in, EvalLib), " ")
			if str != tt.out {
				t.Errorf("got %s, want %s", str, tt.out)
			}
		})
	}
}
