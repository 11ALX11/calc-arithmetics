package app_oop

import (
	"testing"

	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

const (
	TestUnzipOopSuite_expected = TestReadinOopSuite_expected
	TestUnzipOopSuite_filepath = "testdata/test-in.zip"
)

type UnzipOopSuite struct {
	suite.Suite
}

func (s *UnzipOopSuite) BeforeEach(t provider.T) {
	t.Epic("AppOop")
	t.Feature("Input")
	t.Tags("app", "oop", "input", "unzip", "archive", "zip", "reader", "decorator")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("CodeSignal learn course", "https://codesignal.com/learn/courses/large-data-handling-techniques-in-go/lessons/reading-data-from-archived-files-in-go"))
}

func (s *UnzipOopSuite) TestDecryptReadFile(t provider.T) {
	t.Title("Test Unzip ReaderDecorator")
	t.Description("Test ReadFile() in a Unzip.")

	reader := NewReadin()
	file := "../" + TestUnzipOopSuite_filepath // relative to 'app-oop' package
	dataFileInArchive := app.DataFileInArchive

	t.NewStep(
		"Create unzip.",
		allure.NewParameters(
			"reader", reader,
			"dataFileInArchive", dataFileInArchive,
		)...,
	)
	unzip := NewUnzip(reader, dataFileInArchive)

	t.NewStep(
		"Read file using unzip reader.",
		allure.NewParameter(
			"file", file,
		),
	)
	content, err := unzip.
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

	expected := TestUnzipOopSuite_expected
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

func TestUnzipOopSuite(t *testing.T) {
	suite.RunSuite(t, new(UnzipOopSuite))
}
