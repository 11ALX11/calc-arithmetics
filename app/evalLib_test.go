package app

import (
	"fmt"
	"testing"
)

func TestEvalLib(t *testing.T) {
	for _, tt := range evalTests {
		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()
			num := EvalLib(tt.in)
			if fmt.Sprint(num) != tt.out {
				t.Errorf("got %d, want %s", num, tt.out)
			}
		})
	}
}
