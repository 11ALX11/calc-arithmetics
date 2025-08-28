package app

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type EvalLibSuite struct {
	suite.Suite
}

func (s *EvalLibSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Eval")
	t.Tags("app", "math", "lib")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("expr-lang lib", "https://github.com/expr-lang/expr"))
}

func (s *EvalLibSuite) TestEvalLib(t provider.T) {
	t.Title("Test EvalLib()")
	t.Description("Test EvalLib() on a series of strings that contain arithmetic expression")
	t.Parallel()

	for _, tt := range evalTests {
		t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) {
			num := EvalLib(tt.in)
			sCtx.Assert().Equal(tt.out, fmt.Sprint(num), "expected %s, got %s", tt.out, fmt.Sprint(num))
		})
	}
}

func (s *EvalLibSuite) TestEvalLibPairedWithTestsFromFilter(t provider.T) {
	t.Title("Test EvalLib() with tests for filters")
	t.Description("Test EvalLib() with tests for filters paired with ReplaceMathExpressions()")

	var tests []struct{ in, out string }

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	t.Parallel()

	for _, tt := range tests {
		tt := tt // Rebind tt before using it inside the async step.
		t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) {
			str := strings.Trim(ReplaceMathExpressions(tt.in, EvalLib), " ")
			sCtx.Assert().Equal(tt.out, str, "expected %s, got %s", tt.out, str)
		})
	}
}

func TestEvalLibSuite(t *testing.T) {
	suite.RunSuite(t, new(EvalLibSuite))
}
