package app

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ReadZipFileSuite struct {
	suite.Suite
}

func (s *ReadZipFileSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Input")
	t.Tags("app", "input", "archive", "zip")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("CodeSignal learn course", "https://codesignal.com/learn/courses/large-data-handling-techniques-in-go/lessons/reading-data-from-archived-files-in-go"))
}

func (s *ReadZipFileSuite) TestReadZipFile(t provider.T) {
	t.Title("Test zip deciphering")
	t.Description("Test ReadZipFile() on a zip archive containing data.txt")

	expectedString := "asfegf 124 tg ewrhy\n wafdafag wegtwetg 35t\n"
	t.NewStep("Get zip file.", allure.NewParameter("archived string", expectedString))
	// ¯\_(ツ)_/¯

	t.NewStep("Unzip archive and read data.txt's contents.")
	content, err := ReadZipFile("some file", DataFileInArchive)

	t.WithNewStep("Check if there's any error", func(sCtx provider.StepCtx) {
		sCtx.Assert().NoError(err, "Expect no error (nil).")
	}, allure.NewParameter("err", err))

	t.WithNewStep("Compare expected and actual strings.", func(sCtx provider.StepCtx) {
		sCtx.Assert().Equal(expectedString, content, "Expect strings to match.")
	}, allure.NewParameters("expected", expectedString, "actual", content)...)
}

func TestReadZipFileSuite(t *testing.T) {
	suite.RunSuite(t, new(ReadZipFileSuite))
}
