package app

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

const (
	TestDecryptSuite_ciphertext      = "lkoyKPU3Q7YCrDpoIxv6QUNRIyTgU4MSZKDQoF/trX04/2cs3xdaPQTAv8cVg2MAofM23oKMyd8vGTzw7Qp5vO2DqAvCZE5uuyM5emnIwlq/EnoE"
	TestDecryptSuite_expected        = TestEncryptSuite_plaintext
	TestDecryptSuite_key             = TestEncryptSuite_16bit_aes_key
	TestDecryptSuite_wrong_key       = "012345abcdef6789" // gitleaks:allow test vector (not a secret)
	TestDecryptSuite_malformedBase64 = "$$$not_base64$$$"
)

type DecryptSuite struct {
	suite.Suite
}

func (s *DecryptSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Decrypt")
	t.Tags("app", "decrypt", "crypt", "crypto")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("pkg crypto", "https://pkg.go.dev/crypto"))
}

func (s *DecryptSuite) TestDecrypt(t provider.T) {
	t.Title("Test decryption")
	t.Description("Test Decrypt() on a preencoded string with a 16-byte AES key.")

	t.NewStep(
		"Try to decrypt ciphertext with a key.",
		allure.NewParameters(
			"ciphertext", TestDecryptSuite_ciphertext,
			"key", TestDecryptSuite_key,
		)...,
	)

	decodedMsg, err := Decrypt(TestDecryptSuite_ciphertext, TestDecryptSuite_key)

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
			sCtx.Assert().Equal(TestDecryptSuite_expected, decodedMsg, "Expect strings to match.")
		},
		allure.NewParameters(
			"Expected", TestDecryptSuite_expected,
			"Actual", decodedMsg,
		)...,
	)
}

func (s *DecryptSuite) TestDecryptWithWrongKey(t provider.T) {
	t.Title("Test decryption with wrong key")
	t.Description("Test Decrypt() on a preencoded string with a wrong 16-byte AES key and expect an error.")
	t.Severity(allure.CRITICAL) // security issue

	t.NewStep(
		"Try to decrypt ciphertext with wrong key.",
		allure.NewParameters(
			"ciphertext", TestDecryptSuite_ciphertext,
			"key", TestDecryptSuite_wrong_key,
		)...,
	)

	decodedMsg, err := Decrypt(TestDecryptSuite_ciphertext, TestDecryptSuite_wrong_key)

	t.WithNewStep(
		"Check for errors while decrypting.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().EqualError(err, "cipher: message authentication failed", "Expect error")
		},
		allure.NewParameters(
			"Decoded message", decodedMsg,
			"Error", err,
		)...,
	)
}

func (s *DecryptSuite) TestDecryptWithEmptyInput(t provider.T) {
	t.Title("Decrypt with empty input")

	ciphertext := ""

	t.NewStep(
		"Try to decrypt with empty input.",
		allure.NewParameters(
			"ciphertext", ciphertext,
			"key", TestDecryptSuite_key,
		)...,
	)

	_, err := Decrypt(ciphertext, TestDecryptSuite_key)

	t.WithNewStep(
		"Expect error on empty input",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().Error(err)
		},
		allure.NewParameter(
			"Error", err,
		),
	)
}

func (s *DecryptSuite) TestDecryptWithMalformedBase64(t provider.T) {
	t.Title("Decrypt with malformed base64")

	t.NewStep(
		"Try to decrypt with malformed base64.",
		allure.NewParameters(
			"ciphertext", TestDecryptSuite_malformedBase64,
			"key", TestDecryptSuite_key,
		)...,
	)

	_, err := Decrypt(TestDecryptSuite_malformedBase64, TestDecryptSuite_key)

	t.WithNewStep(
		"Expect base64 decode error",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().Error(err)
		},
		allure.NewParameter(
			"Error", err,
		),
	)
}

func TestDecryptSuite(t *testing.T) {
	suite.RunSuite(t, new(DecryptSuite))
}
