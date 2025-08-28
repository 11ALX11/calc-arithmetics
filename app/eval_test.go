package app

import (
	"fmt"
	"strings"
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
	{"+ - + - - (--1)", "-1"},
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
	t.Title("Test Eval()")
	t.Description("Test Eval() on a series of strings that contain arithmetic expression")
	t.Parallel()

	for _, tt := range evalTests {
		t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) {
			num := Eval(tt.in)
			sCtx.Assert().Equal(tt.out, fmt.Sprint(num), "expected %s, got %s", tt.out, fmt.Sprint(num))
		})
	}
}

func (s *EvalSuite) TestEvalPairedWithTestsFromFilter(t provider.T) {
	t.Title("Test Eval() with tests for filters")
	t.Description("Test Eval() with tests for filters paired with ReplaceMathExpressions()")

	var tests []struct{ in, out string }

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	t.Parallel()

	for _, tt := range tests {
		tt := tt // Rebind tt before using it inside the async step.
		t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) {
			str := strings.Trim(ReplaceMathExpressions(tt.in, Eval), " ")
			sCtx.Assert().Equal(tt.out, str, "expected %s, got %s", tt.out, str)
		})
	}
}

func TestEvalSuite(t *testing.T) {
	suite.RunSuite(t, new(EvalSuite))
}
