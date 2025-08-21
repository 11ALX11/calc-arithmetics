package app

import (
	"strings"
	"testing"

	"github.com/GRbit/go-pcre"
	"github.com/ozontech/allure-go/pkg/allure"
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

	var tests []struct{ in, out string }

	tests = append(tests, filterTests...)
	tests = append(tests, evalTests...)

	// t.Parallel() //ToDo: uncomment after debug

	for _, tt := range tests {
		// t.WithNewAsyncStep(tt.in, func(sCtx provider.StepCtx) { //ToDo: uncomment after debug
		t.WithNewStep(tt.in, func(sCtx provider.StepCtx) { //ToDo: delete line after debug
			str := strings.Trim(ReplaceMathExpressionsRegex(tt.in, EvalLib), " ")
			sCtx.Assert().Equal(tt.out, str, "expected %s, got %s", tt.out, str)
		})
	}
}

func (s *FilterRegexSuite) TestRegexLib(t provider.T) {
	t.Description("Test github.com/GRbit/go-pcre lib with simple pattern and string")

	const simplePattern = `\d+`
	t.NewStep("Define pattern, compile it and get matcher object.", allure.NewParameter("pattern", simplePattern))

	var simpleRe = pcre.MustCompileJIT(simplePattern, 0, pcre.CONFIG_JIT)
	var testString = "abc 123 def"
	var simpleMatcher = *simpleRe.NewMatcherString(testString, 0)

	t.NewStep("Try to get a match from a test string and verify it.", allure.NewParameter("testString", testString))

	if simpleMatcher.Matches {

		match := simpleMatcher.GroupString(0)
		t.Assert().Equal("123", match, "Expect a match to be '123'")

	} else {
		t.Errorf("simpleMatcher.Matches is false. Last match wasn't successful.")
	}
}

func TestFilterRegexSuite(t *testing.T) {
	suite.RunSuite(t, new(FilterRegexSuite))
}
