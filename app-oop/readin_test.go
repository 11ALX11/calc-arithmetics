package app_oop

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

const (
	TestReadinOopSuite_expected = "asfegf 124 tg ewrhy\n wafdafag wegtwetg 35t\n"
	TestReadinOopSuite_filepath = "testdata/test-in.txt"
)

type ReadinOopSuite struct {
	suite.Suite
}

func (s *ReadinOopSuite) BeforeEach(t provider.T) {
	t.Epic("AppOop")
	t.Feature("Input")
	t.Tags("app", "oop", "input", "reader")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
}

func (s *ReadinOopSuite) TestReadFile(t provider.T) {
	t.Title("Test Readin Reader")
	t.Description("Test ReadFile() in a Readin.")

	file := "../" + TestReadinOopSuite_filepath // relative to 'app-oop' package

	t.NewStep("Create reader.")
	reader := NewReadin()

	t.NewStep(
		"Read file using reader.",
		allure.NewParameter(
			"file", file,
		),
	)
	content, err := reader.
		ReadFile(file).
		GetContentError()

	t.WithNewStep(
		"Check error.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error")
		},
		allure.NewParameter(
			"err", err,
		),
	)

	expected := TestReadinOopSuite_expected
	t.WithNewStep(
		"Compare strings, expect to match.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().Equal(expected, content, "Expect to match")
		},
		allure.NewParameters(
			"expected", expected,
			"actual", content,
		)...,
	)
}

func TestReadinOopSuite(t *testing.T) {
	suite.RunSuite(t, new(ReadinOopSuite))
}
