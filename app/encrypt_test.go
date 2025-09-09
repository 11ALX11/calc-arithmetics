package app

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

const (
	TestEncryptSuite_16bit_aes_key = "0123456789abcdef" // gitleaks:allow test vector (not a secret)
	TestEncryptSuite_plaintext     = "woehfoew fewoifh ewiofhwpfqf 	owgh	obe	wgou	ewhgjbngd gd"
)

type EncryptSuite struct {
	suite.Suite
}

func (s *EncryptSuite) BeforeEach(t provider.T) {
	t.Epic("App")
	t.Feature("Encrypt")
	t.Tags("app", "encrypt", "crypt", "crypto")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("pkg crypto", "https://pkg.go.dev/crypto"))
}

func (s *EncryptSuite) TestEncrypt(t provider.T) {
	t.Title("Test encryption")
	t.Description("Test Encrypt() on a string with a 16-bit AES key. Uses Decrypt() to decrypt back for comparison with original.")

	t.NewStep(
		"Try to encrypt plaintext with a key.",
		allure.NewParameters(
			"plaintext", TestEncryptSuite_plaintext,
			"key", TestEncryptSuite_16bit_aes_key,
		)...,
	)
	encodedMsg, err := Encrypt(TestEncryptSuite_plaintext, TestEncryptSuite_16bit_aes_key)

	t.WithNewStep(
		"Check for errors while encrypting.",
		func(sCtx provider.StepCtx) {
			sCtx.Assert().NoError(err, "Expect no error (nil)")
		},
		allure.NewParameters(
			"Encoded Message", encodedMsg,
			"Error", err,
		)...,
	)

	decodedMsg, err := Decrypt(encodedMsg, TestEncryptSuite_16bit_aes_key)

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
			sCtx.Assert().Equal(TestEncryptSuite_plaintext, decodedMsg, "Expect strings to match.")
		},
		allure.NewParameters(
			"Expected", TestEncryptSuite_plaintext,
			"Actual", decodedMsg,
		)...,
	)
}

func TestEncryptSuite(t *testing.T) {
	suite.RunSuite(t, new(EncryptSuite))
}
