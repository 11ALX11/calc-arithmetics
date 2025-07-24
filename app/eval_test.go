package app

import (
	"fmt"
	"testing"
)

var evalTests = []struct {
	in  string
	out string
}{
	{"1 + 1", "2"},
	{"6-4 / 2 ", "4"},
	{"2*(5+5*2)/3+(6/2+8)", "21"},
	{"(2+6* 3+5- (3*14/7+2)*5)+3", "-12"},
}

func TestEval(t *testing.T) {
	for _, tt := range evalTests {
		t.Run(tt.in, func(t *testing.T) {
			num := Eval(tt.in)
			if fmt.Sprint(num) != tt.out {
				t.Errorf("got %d, want %s", num, tt.out)
			}
		})
	}
}
