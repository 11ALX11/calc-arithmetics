package app

import (
	"strings"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type FilterRegexSuite struct {
	suite.Suite
}

func (s *FilterRegexSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Filter")
	t.Tags("app", "math", "regex")
}

func (s *FilterRegexSuite) TestFilterRegex(t provider.T) {
	t.Description("Test ReplaceMathExpressionsRegex() on a series of strings that contain arithmetic expression to filter from sentences (using EvalLib())")

	t.XSkip()
	t.Skip()

	var tests []struct{ in, out string }

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	for _, tt := range tests {
		t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) {
			str := strings.Trim(ReplaceMathExpressionsRegex(tt.in, EvalLib), " ")
			sCtx.Assert().Equal(tt.out, str, "expected %s, got %s", tt.out, str)
		})
	}
}

func TestFilterRegexSuite(t *testing.T) {
	suite.RunSuite(t, new(FilterRegexSuite))
}
