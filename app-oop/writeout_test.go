package app_oop

import (
	"os"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type WriteoutOopSuite struct {
	suite.Suite
}

const (
	WriteoutOop_filepath = "tmp-test-out.txt"
	WriteoutOop_text     = "asfegf 124 tg ewrhy\n wafdafag wegtwetg 35t\n"
)

func (s *WriteoutOopSuite) BeforeEach(t provider.T) {
	t.Epic("AppOop")
	t.Feature("Output")
	t.Tags("app", "oop", "output", "writer")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
}

func (s *WriteoutOopSuite) TestWriteFile(t provider.T) {
	t.Title("Test Writeout Writer")
	t.Description("Test WriteFile() in Writeout.")

	var file string
	t.WithNewStep(
		"Create tmp files.",
		func(sCtx provider.StepCtx) {
			file = createTmpFile(sCtx, WriteoutOop_filepath, "")
		},
	)

	expectedString := WriteoutOop_text
	t.NewStep(
		"Try to write to a txt file.",
		allure.NewParameters(
			"String to write", expectedString,
			"File", file,
		)...,
	)

	err := NewWriteout().
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

	bytes, err := os.ReadFile(file)
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

func TestWriteoutOopSuite(t *testing.T) {
	suite.RunSuite(t, new(WriteoutOopSuite))
}
