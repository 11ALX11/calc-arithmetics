package app

import (
	"fmt"
	"testing"

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
}

func (s *EvalLibSuite) TestEvalLib(t provider.T) {
	t.Parallel()
	for _, tt := range evalTests {
		t.Run(tt.in, func(t provider.T) {
			// ToDo: BeforeEach somehow doesnt apply
			t.Epic("App")
			t.Feature("Eval")
			t.Tags("app", "math", "lib", "parallel")

			tti := tt
			t.Parallel()

			num := EvalLib(tti.in)
			t.Assert().Equal(fmt.Sprint(num), tti.out, "expected %s, got %s", tti.out, fmt.Sprint(num))
		})
	}
}

func TestEvalLibSuite(t *testing.T) {
	suite.RunSuite(t, new(EvalLibSuite))
}
