package app

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

const (
	Test_in_txt_content  = "asfegf 124 tg ewrhy\n wafdafag wegtwetg 35t\n"
	Test_in_txt_filepath = "testdata/test-in.txt"
)

type ReadFileSuite struct {
	suite.Suite
}

func (s *ReadFileSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Input")
	t.Tags("app", "input")
	t.Severity(allure.CRITICAL)
	t.Owner("github.com/11ALX11")
}

func (s *ReadFileSuite) TestReadFile(t provider.T) {
	t.Title("Test file reading")
	t.Description("Test ReadFile() on a txt file.")

	expectedString := Test_in_txt_content
	file := "../" + Test_in_txt_filepath // relative to 'app' package
	t.NewStep(
		"Try to get txt file.",
		allure.NewParameters(
			"String in a file", expectedString,
			"File", file,
		)...,
	)

	t.NewStep("Try to read file's contents.")
	content, err := ReadFile(file)

	t.WithNewStep("Check if there's any error", func(sCtx provider.StepCtx) {
		sCtx.Assert().NoError(err, "Expect no error (nil).")
	}, allure.NewParameter("Error", err))

	t.WithNewStep(
		"Compare expected and actual strings.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().Equal(expectedString, content, "Expect strings to match.")
		},
		allure.NewParameters(
			"Expected", expectedString,
			"Actual", content,
		)...,
	)
}

func TestReadFileSuite(t *testing.T) {
	suite.RunSuite(t, new(ReadFileSuite))
}
