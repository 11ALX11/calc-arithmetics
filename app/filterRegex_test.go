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
	var tests []struct{ in, out string }

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	for _, tt := range tests {
		t.Run(tt.in, func(t provider.T) {
			// ToDo: BeforeEach somehow doesnt apply
			t.Epic("App")
			t.Feature("Filter")
			t.Tags("app", "math", "regex", "parallel")

			tti := tt
			t.Parallel()

			str := strings.Trim(ReplaceMathExpressionsRegex(tti.in, EvalLib), " ")
			t.Assert().Equal(str, tti.out, "expected %s, got %s", tti.out, str)
		})
	}
}

func TestFilterRegexSuite(t *testing.T) {
	suite.RunSuite(t, new(FilterRegexSuite))
}
