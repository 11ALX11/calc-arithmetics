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

const Test_in_zip_content = Test_in_txt_content
const Test_in_zip_filepath = "testdata/test-in.zip"

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

	expectedString := Test_in_zip_content
	file := "../" + Test_in_zip_filepath // relative to 'app' package
	t.NewStep("Try to get zip file.", allure.NewParameters("Archived string", expectedString, "File", file)...)

	t.NewStep("Try to unzip archive and read data.txt's contents.")
	content, err := ReadZipFile(file, DataFileInArchive)

	t.WithNewStep("Check if there's any error", func(sCtx provider.StepCtx) {
		sCtx.Assert().NoError(err, "Expect no error (nil).")
	}, allure.NewParameter("Error", err))

	t.WithNewStep("Compare expected and actual strings.", func(sCtx provider.StepCtx) {
		sCtx.Assert().Equal(expectedString, content, "Expect strings to match.")
	}, allure.NewParameters("Expected", expectedString, "Actual", content)...)
}

func TestReadZipFileSuite(t *testing.T) {
	suite.RunSuite(t, new(ReadZipFileSuite))
}
