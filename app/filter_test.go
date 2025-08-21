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
	{"Here's another one: (2+6* 3+5- (3*14/(7+2)*5)+3 = 4.667.", "Here's another one: (4.667 = 4.667."},
	{"Here's another one: 2+6* 3+5- (3*14/(7+2)*5)+3 = 4.667.", "Here's another one: 4.667 = 4.667."},
	{"Here's another one: 2+6* 3+5- (3*14(/7+2)*5)+3", "Here's another one: 25- (42(/9)*5)3"},
	{"Here's multiples: (2+6* 3+5- (3*14/7+2)*5)+3 = -12, 1 + 2 + 3 + 4 - 5 = 5", "Here's multiples: -12 = -12, 5 = 5"},
	{"Here's an arithmetic () (1+1,2+4) 1+1=2.", "Here's an arithmetic () (2,6) 2=2."},
	{"Here's an arithmetic 1+d1=2.", "Here's an arithmetic 1+d1=2."},
	{"Here's an arithmetic (1+)d1=2.", "Here's an arithmetic (1+)d1=2."},
	{"Here's an arithmetic 1+1d=2.", "Here's an arithmetic 2d=2."},
	{"Here's an arithmetic (1+(1))=2.", "Here's an arithmetic 2=2."},
	{"Here's an arithmetic (1+(-1))=0.", "Here's an arithmetic 0=0."},
	{"Here's an arithmetic (1+-1)=0.", "Here's an arithmetic 0=0."},
	{"Here's an arithmetic (1+(-1)=0.", "Here's an arithmetic (0=0."},
	{"Here's an arithmetic (1+-1))=0.", "Here's an arithmetic 0)=0."},
	{"Here's an arithmetic (1+)-1)=0.", "Here's an arithmetic (1+)-1)=0."},
	{"Here's an arithmetic )1+1(=2.", "Here's an arithmetic )2(=2."},
	{"Here's an arithmetic (1+1)+1+1)", "Here's an arithmetic 4)"},
	{"Here's an arithmetic (1+1+(1+1)", "Here's an arithmetic (4"},
	{"Here's an arithmetic (1+1+(1+1", "Here's an arithmetic (2+(2"},
	{"Here's an arithmetic 1+1)+1+1)", "Here's an arithmetic 2)2)"},
	{"Here's no digits ().", "Here's no digits ()."},
	{"Here's no digits (+).", "Here's no digits (+)."},
	{"Here's a digit (+1).", "Here's a digit 1."},
	{"Here's a digit (-1).", "Here's a digit -1."},
	{"Here's a digit (1+).", "Here's a digit (1+)."},
	{"t -(-1)", "t 1"},
	{"t --1", "t 1"},
	{"t ---1", "t -1"},
	{"t -+-1", "t 1"},
	{"t 1-+-1", "t 2"},
	{"t 1**1", "t 1**1"},
	{"t 1-/1", "t 1-/1"},
	{"t 1+*1", "t 1+*1"},
	{"*1", "*1"},
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
	t.Description("Test ReplaceMathExpressions() on a series of strings that contain arithmetic expression to filter from sentences (using EvalLib())")

	var tests []struct{ in, out string }

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	t.Parallel()

	for _, tt := range tests {
		t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) {
			str := strings.Trim(ReplaceMathExpressions(tt.in, EvalLib), " ")
			sCtx.Assert().Equal(tt.out, str, "expected %s, got %s", tt.out, str)
		})
	}
}

func TestFilterSuite(t *testing.T) {
	suite.RunSuite(t, new(FilterSuite))
}
