package app

import (
	"strings"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

var filterTests = []struct {
	in  string
	out string
}{
	{"Here's an arithmetic 1+1=2.", "Here's an arithmetic 2=2."},
	{"Here's another one: (2+6* 3+5- (3*14/7+2)*5)+3 = -12.", "Here's another one: -12 = -12."},
	{"Here's multiples: (2+6* 3+5- (3*14/7+2)*5)+3 = -12, 1 + 2 + 3 + 4 - 5 = 5", "Here's multiples: -12 = -12, 5 = 5"},
	{"Here's an arithmetic (1+1,2+4) 1+1=2.", "Here's an arithmetic (2,6) 2=2."},
}

type FilterSuite struct {
	suite.Suite
}

func (s *FilterSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Filter")
	t.Tags("app", "math")
}

func (s *FilterSuite) TestFilter(t provider.T) {
	var tests []struct{ in, out string }

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	t.Parallel()

	for _, tt := range tests {
		t.Run(tt.in, func(t provider.T) {
			// ToDo: BeforeEach somehow doesnt apply
			t.Epic("App")
			t.Feature("Filter")
			t.Tags("app", "math", "parallel")

			tti := tt
			t.Parallel()

			str := strings.Trim(ReplaceMathExpressions(tti.in, EvalLib), " ")
			t.Assert().Equal(str, tti.out, "expected %s, got %s", tti.out, str)
		})
	}
}

func TestFilterSuite(t *testing.T) {
	suite.RunSuite(t, new(FilterSuite))
}
