package app

import (
	"strings"
	"testing"
)

func TestFilterRegex(t *testing.T) {

	var tests []struct {
		in  string
		out string
	}

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			str := strings.Trim(ReplaceMathExpressionsRegex(tt.in, EvalLib), " ")
			if str != tt.out {
				t.Errorf("got %s, want %s", str, tt.out)
			}
		})
	}
}
