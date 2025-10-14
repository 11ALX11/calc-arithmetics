package app

import (
	"os"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type WriteZipFileSuite struct {
	suite.Suite
}

const Tmp_Test_out_zip_filepath = "tmp-test-out.zip"

func (s *WriteZipFileSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Output")
	t.Tags("app", "output", "archive", "zip")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("stackoverflow source of info", "https://stackoverflow.com/questions/37869793/how-do-i-zip-a-directory-containing-sub-directories-or-files-in-golang"))
}

func (s *WriteZipFileSuite) TestWriteZipFile(t provider.T) {
	t.Title("Test zip archiving")
	t.Description("Test WriteZipFile() on a test string. Check is done with ReadZipFile().")

	file, err := os.CreateTemp("", Tmp_Test_out_zip_filepath)
	defer os.Remove(file.Name())

	t.WithNewStep(
		"Try to create temporary file",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameters(
			"File", Tmp_Test_out_zip_filepath,
			"TmpFile", file.Name(),
		)...,
	)

	expectedString := Test_in_txt_content // Any string will do.
	t.NewStep(
		"Try to create a zip file.",
		allure.NewParameters(
			"String to write", expectedString,
			"File", file.Name(),
		)...,
	)

	err = WriteZipFile(file.Name(), expectedString, DataFileInArchive)

	t.WithNewStep(
		"Check if there's any error",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameter(
			"Error", err,
		),
	)

	content, err := ReadZipFile(file.Name(), DataFileInArchive)

	t.WithNewStep(
		"Check if there's any error while reading zip file",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameter(
			"Error", err,
		),
	)

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

func (s *WriteZipFileSuite) TestGetZipData(t provider.T) {
	t.Title("Test zip archiving (GetZipData())")
	t.Description("Test GetZipData() on a test string. Check is done with ReadZipData().")
	t.Link(allure.LinkLink("Example code", "https://www.programmersought.com/article/28081546124/"))

	expectedString := Test_in_txt_content // Any string will do.

	t.NewStep(
		"Try to get zip data.",
		allure.NewParameter(
			"String to write", expectedString,
		),
	)

	archived, err := GetZipData(expectedString, DataFileInArchive)

	t.WithNewStep(
		"Check if there's any error",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameters(
			"Data", archived,
			"Error", err,
		)...,
	)

	content, err := ReadZipData(archived, DataFileInArchive)

	t.WithNewStep(
		"Check if there's any error while reading zip data",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameter(
			"Error", err,
		),
	)

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

func TestWriteZipFileSuite(t *testing.T) {
	suite.RunSuite(t, new(WriteZipFileSuite))
}
