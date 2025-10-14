package app_oop

import (
	"testing"

	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type ArchiveOopSuite struct {
	suite.Suite
}

const (
	ArchiveOop_filepath = "tmp-test-out.zip"
	ArchiveOop_text     = WriteoutOop_text
)

func (s *ArchiveOopSuite) BeforeEach(t provider.T) {
	t.Epic("AppOop")
	t.Feature("Output")
	t.Tags("app", "oop", "output", "archive", "zip", "writer", "decorator")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("stackoverflow source of info", "https://stackoverflow.com/questions/37869793/how-do-i-zip-a-directory-containing-sub-directories-or-files-in-golang"))
}

func (s *ArchiveOopSuite) TestWriteFile(t provider.T) {
	t.Title("Test Archive WriterDecorator")
	t.Description("Test WriteFile() in Archive. Uses Unzip Reader to check file contents.")

	var file string
	t.WithNewStep(
		"Create tmp files.",
		func(sCtx provider.StepCtx) {
			file = createTmpFile(sCtx, ArchiveOop_filepath, "")
		},
	)

	expectedString := ArchiveOop_text // Any string will do.
	t.NewStep(
		"Try to create a zip file.",
		allure.NewParameters(
			"String to write", expectedString,
			"File", file,
		)...,
	)

	writer := NewWriteout()

	err := NewArchive(writer, app.DataFileInArchive).
		WriteFile(file, expectedString).
		GetError()

	t.WithNewStep(
		"Check if there's any error",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameter(
			"Error", err,
		),
	)

	reader := NewReadin()
	content, err := NewUnzip(reader, app.DataFileInArchive).
		ReadFile(file).
		GetContentError()

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

func TestArchiveOopSuite(t *testing.T) {
	suite.RunSuite(t, new(ArchiveOopSuite))
}
