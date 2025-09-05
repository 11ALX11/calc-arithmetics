package app

import (
	"os"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type WriteFileSuite struct {
	suite.Suite
}

func (s *WriteFileSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Output")
	t.Tags("app", "output")
	t.Severity(allure.CRITICAL)
	t.Owner("github.com/11ALX11")
}

func (s *WriteFileSuite) TestReadFile(t provider.T) {
	t.Title("Test file writing")
	t.Description("Test WriteFile() on a txt file.")

	expectedString := "asfegf 124 tg ewrhy\n wafdafag wegtwetg 35t\n"
	file := "./tmp/some-file.txt"
	t.NewStep("Write txt file.", allure.NewParameters("String to write", expectedString, "File", file)...)

	// ¯\_(ツ)_/¯
	err := WriteFile(file, expectedString)

	t.WithNewStep("Check if there's any error", func(sCtx provider.StepCtx) {
		sCtx.Assert().NoError(err, "Expect no error (nil).")
	}, allure.NewParameter("Error", err))

	content, err := os.ReadFile(file)

	t.WithNewStep("Check if there's any error while reading file", func(sCtx provider.StepCtx) {
		sCtx.Assert().NoError(err, "Expect no error (nil).")
	}, allure.NewParameter("Error", err))

	t.WithNewStep("Compare expected and actual strings.", func(sCtx provider.StepCtx) {
		sCtx.Assert().Equal(expectedString, content, "Expect strings to match.")
	}, allure.NewParameters("Expected", expectedString, "Actual", content)...)
}

func TestWriteFileSuite(t *testing.T) {
	suite.RunSuite(t, new(WriteFileSuite))
}
