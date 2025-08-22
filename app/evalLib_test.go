package app

import (
	"fmt"
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

func TestEvalLibSuite(t *testing.T) {
	suite.RunSuite(t, new(EvalLibSuite))
}
