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

const Tmp_Test_out_txt_filepath = "tmp-test-out.txt"

func (s *WriteFileSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Output")
	t.Tags("app", "output")
	t.Severity(allure.CRITICAL)
	t.Owner("github.com/11ALX11")
}

func (s *WriteFileSuite) TestWriteFile(t provider.T) {
	t.Title("Test file writing")
	t.Description("Test WriteFile() on a txt file.")

	file, err := os.CreateTemp("", Tmp_Test_out_txt_filepath)
	defer os.Remove(file.Name())

	t.WithNewStep(
		"Try to create temporary file",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameters(
			"File", Tmp_Test_out_txt_filepath,
			"TmpFile", file.Name(),
		)...,
	)

	expectedString := Test_in_txt_content // Any string will do.
	t.NewStep(
		"Try to write to a txt file.",
		allure.NewParameters(
			"String to write", expectedString,
			"File", file.Name(),
		)...,
	)

	err = WriteFile(file.Name(), expectedString)

	t.WithNewStep(
		"Check if there's any error",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameter(
			"Error", err,
		),
	)

	bytes, err := os.ReadFile(file.Name())
	content := string(bytes)

	t.WithNewStep(
		"Check if there's any error while reading file",
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

func TestWriteFileSuite(t *testing.T) {
	suite.RunSuite(t, new(WriteFileSuite))
}
