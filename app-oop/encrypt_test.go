package app_oop

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

const (
	TestEncryptOopSuite_key         = TestDecryptOopSuite_key
	TestEncryptOopSuite_keyPath     = TestDecryptOopSuite_keyPath
	TestEncryptOopSuite_plaintext   = TestDecryptOopSuite_expected
	TestEncryptOopSuite_tmpFilePath = "tmp-out.aes"
)

type EncryptOopSuite struct {
	suite.Suite
}

func (s *EncryptOopSuite) BeforeEach(t provider.T) {
	t.Epic("AppOop")
	t.Feature("Output")
	t.Tags("app", "oop", "output", "encrypt", "crypt", "crypto", "writer", "decorator")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("pkg crypto", "https://pkg.go.dev/crypto"))
}

func (s *EncryptOopSuite) TestWriteFile(t provider.T) {
	t.Title("Test Encrypt WriterDecorator")
	t.Description("Test WriteFile() in Encrypt. Uses Decrypt Reader to decrypt file contents.")

	var keyPath, file string
	t.WithNewStep(
		"Create tmp files.",
		func(sCtx provider.StepCtx) {
			keyPath = createTmpFile(sCtx, TestEncryptOopSuite_keyPath, TestEncryptOopSuite_key)
			file = createTmpFile(sCtx, TestEncryptOopSuite_tmpFilePath, "")
		},
	)

	t.NewStep(
		"Try to encrypt plaintext with a key.",
		allure.NewParameters(
			"plaintext", TestEncryptOopSuite_plaintext,
			"key", TestEncryptOopSuite_key,
		)...,
	)
	writer := NewWriteout()
	err := NewEncrypt(writer, keyPath).
		WriteFile(file, TestEncryptOopSuite_plaintext).
		GetError()

	t.WithNewStep(
		"Check for errors while encrypting.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil)")
		},
		allure.NewParameters(
			"Error", err,
		)...,
	)

	reader := NewReadin()
	decodedMsg, err := NewDecrypt(reader, keyPath).
		ReadFile(file).
		GetContentError()

	t.WithNewStep(
		"Check for errors while decrypting.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil)")
		},
		allure.NewParameter(
			"Error", err,
		),
	)

	t.WithNewStep(
		"Compare expected and actual strings.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().Equal(TestEncryptOopSuite_plaintext, decodedMsg, "Expect strings to match.")
		},
		allure.NewParameters(
			"Expected", TestEncryptOopSuite_plaintext,
			"Actual", decodedMsg,
		)...,
	)
}

func TestEncryptOopSuite(t *testing.T) {
	suite.RunSuite(t, new(EncryptOopSuite))
}
