package app

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
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

type EvalSuite struct {
	suite.Suite
}

func (s *EvalSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Eval")
	t.Tags("app", "math")
}

func (s *EvalSuite) TestEval(t provider.T) {
	t.Parallel()
	for _, tt := range evalTests {
		t.Run(tt.in, func(t provider.T) {
			// ToDo: BeforeEach somehow doesnt apply
			t.Epic("App")
			t.Feature("Eval")
			t.Tags("app", "math", "parallel")

			tti := tt
			t.Parallel()

			num := Eval(tti.in)
			t.Assert().Equal(fmt.Sprint(num), tti.out, "expected %s, got %s", tti.out, fmt.Sprint(num))
		})
	}
}

func TestEvalSuite(t *testing.T) {
	suite.RunSuite(t, new(EvalSuite))
}
