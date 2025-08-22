package app

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

var evalTests = []struct {
	in  string
	out string
}{
	{"1 + (1)", "2"},
	{"6-4 / 2 ", "4"},
	{"2*(5+5*2)/3+(6/2+8)", "21"},
	{"(2+6* 3+5- (3*14/7+2)*5)+3", "-12"},
	{"1 + -1", "0"},
	{"1 + -1*(-1)", "2"},
	{"1 + -1 * -1", "2"},
	{"1 - (-1)", "2"},
}

type EvalSuite struct {
	suite.Suite
}

func (s *EvalSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Eval")
	t.Tags("app", "math")
	t.Severity(allure.CRITICAL)
	t.Owner("github.com/11ALX11")
}

func (s *EvalSuite) TestEval(t provider.T) {
	t.XSkip()
	t.Skip()

	t.Title("Test Eval()")
	t.Description("Test Eval() on a series of strings that contain arithmetic expression")
	t.Link(allure.IssueLink("https://github.com/11ALX11/calc-arithmetics/issues/9")) // temporary
	t.Parallel()

	for _, tt := range evalTests {
		t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) {
			num := Eval(tt.in)
			sCtx.Assert().Equal(tt.out, fmt.Sprint(num), "expected %s, got %s", tt.out, fmt.Sprint(num))
		})
	}
}

func TestEvalSuite(t *testing.T) {
	suite.RunSuite(t, new(EvalSuite))
}
