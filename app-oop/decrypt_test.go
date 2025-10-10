package app_oop

import (
	"os"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

const (
	TestDecryptOopSuite_key         = "0123456789abcdef" // gitleaks:allow test vector (not a secret)
	TestDecryptOopSuite_keyPath     = "tmp-key16.txt"
	TestDecryptOopSuite_ciphertext  = "lkoyKPU3Q7YCrDpoIxv6QUNRIyTgU4MSZKDQoF/trX04/2cs3xdaPQTAv8cVg2MAofM23oKMyd8vGTzw7Qp5vO2DqAvCZE5uuyM5emnIwlq/EnoE"
	TestDecryptOopSuite_expected    = "woehfoew fewoifh ewiofhwpfqf 	owgh	obe	wgou	ewhgjbngd gd"
	TestDecryptOopSuite_tmpFilePath = "tmp-in.aes"
)

type DecryptOopSuite struct {
	suite.Suite
}

func (s *DecryptOopSuite) BeforeEach(t provider.T) {
	t.Epic("AppOop")
	t.Feature("Input")
	t.Tags("app", "oop", "input", "decrypt", "crypt", "crypto", "reader", "decorator")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("pkg crypto", "https://pkg.go.dev/crypto"))
}

func createTmpFile(t provider.StepCtx, filePath, content string) string {
	file, err := os.CreateTemp("", filePath)
	// defer os.Remove(file.Name())

	t.WithNewStep(
		"Try to create temporary file",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameters(
			"File", filePath,
			"TmpFile", file.Name(),
		)...,
	)
	if err != nil {
		return ""
	}

	_, err = file.Write([]byte(content))

	t.WithNewStep(
		"Try to write to temporary file",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil).")
		},
		allure.NewParameters(
			"File", file.Name(),
			"Content", content,
		)...,
	)
	if err != nil {
		return ""
	}

	return file.Name()
}

func (s *DecryptOopSuite) TestDecryptReadFile(t provider.T) {
	t.Title("Test Decrypt ReaderDecorator")
	t.Description("Test ReadFile() in a Decrypt.")

	reader := NewReadin()

	var keyPath, file string
	t.WithNewStep(
		"Create tmp files.",
		func(sCtx provider.StepCtx) {
			keyPath = createTmpFile(sCtx, TestDecryptOopSuite_keyPath, TestDecryptOopSuite_key)
			file = createTmpFile(sCtx, TestDecryptOopSuite_tmpFilePath, TestDecryptOopSuite_ciphertext)
		},
	)

	t.NewStep(
		"Create decrypt.",
		allure.NewParameters(
			"reader", reader,
			"key", TestDecryptOopSuite_key,
			"keyPath", keyPath,
		)...,
	)
	decrypt := NewDecrypt(reader, keyPath)

	t.NewStep(
		"Read file using decrypt reader.",
		allure.NewParameter(
			"file", file,
		),
	)
	content, err := decrypt.
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

	expected := TestDecryptOopSuite_expected
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

func TestDecryptOopSuite(t *testing.T) {
	suite.RunSuite(t, new(DecryptOopSuite))
}
