package i18n

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type I18nSuite struct {
	suite.Suite
}

func (s *I18nSuite) BeforeAll(t provider.T) {
	_ = godotenv.Load("../.env")
}

func (s *I18nSuite) BeforeEach(t provider.T) {
	t.Epic("i18n")
	t.Feature("i18n-core")
	t.Tags("i18n", "env")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("PhraseApp-Blog i18n github", "https://github.com/PhraseApp-Blog/go-internationalization/tree/master/pkg/i18n"))
}

func (s *I18nSuite) TestInit(t provider.T) {
	t.Title("Test Init()")
	t.Description("Test Init: loading environment, initializing translations and checking basic translations")

	err := Init()
	t.Assert().NoError(err, "Expected no error on Init()")

	t.Assert().True(len(langMap) >= 2, "Expected at least 2 languages, got %d", len(langMap))

	t.Assert().Equal("Test", EN.T("Test"), "Expected translation to EN to match")
	t.Assert().Equal("Тест", RU.T("Test"), "Expected translation to RU to match")
}

func (s *I18nSuite) getLanguageCodeTestBody(t provider.T, envVarName string) {
	t.Title("Test getLanguageCode()")
	t.Description("Test getLanguageCode with " + envVarName + " env")
	t.WithParameters(allure.NewParameter("env", envVarName))

	t.Setenv(envVarName, "ru_RU")

	lang := getLanguageCode()
	t.Assert().Equal("ru_RU", lang, "Expected language code to be ru_RU")
}

func (s *I18nSuite) TestGetLanguageCodeLANGUAGE(t provider.T) {
	s.getLanguageCodeTestBody(t, "LANGUAGE")
}

func (s *I18nSuite) TestGetLanguageCodeLC_ALL(t provider.T) {
	s.getLanguageCodeTestBody(t, "LC_ALL")
}

func (s *I18nSuite) TestGetLanguageCodeLC_MESSAGES(t provider.T) {
	s.getLanguageCodeTestBody(t, "LC_MESSAGES")
}

func (s *I18nSuite) TestGetLanguageCodeLANG(t provider.T) {
	s.getLanguageCodeTestBody(t, "LANG")
}

func (s *I18nSuite) TestGetSupportedLanguages(t provider.T) {
	t.Title("Test GetSupportedLanguages()")
	t.Description("Test GetSupportedLanguages returns at least 2 languages")

	Init()

	supported := GetSupportedLanguages()
	t.Assert().True(len(supported) >= 2, "Expected at least 2 supported languages, got %d", len(supported))
}

func (s *I18nSuite) TestNewLanguageFromString(t provider.T) {
	t.Title("Test NewLanguageFromString()")
	t.Description("Test NewLanguageFromString parses language codes")

	t.Assert().Equal(NewLanguageFromString("ru_RU"), RU, "Expected ru_RU to be RU")
	t.Assert().Equal(NewLanguageFromString("ru"), RU, "Expected ru to be RU")
	t.Assert().Equal(NewLanguageFromString("en_US"), EN, "Expected en_US to be EN")
}

func (s *I18nSuite) TestGetCurrentLanguage(t provider.T) {
	t.Title("Test GetCurrentLanguage()")
	t.Description("Test GetCurrentLanguage returns correct language based on LANGUAGE env")

	t.Setenv("LANGUAGE", "ru_RU")
	Init()

	lang := GetCurrentLanguage()
	t.Assert().Equal(RU, lang, "Expected current language to be RU")
}

func (s *I18nSuite) TestT(t provider.T) {
	t.Title("Test i18n.T()")
	t.Description("Test T() returns correct translation for RU")

	t.Setenv("LANGUAGE", "ru_RU")
	Init()

	expected := "Тест"
	got := T("Test")
	t.Assert().Equal(expected, got, "Expected translation to RU to match")
}

func (s *I18nSuite) TestSetCurrentLocale(t provider.T) {
	t.Title("Test SetCurrentLocale()")
	t.Description("Test SetCurrentLocale switches languages correctly")

	Init()
	t.Setenv("LANGUAGE", "ru_RU")

	SetCurrentLocale(EN.String())
	t.Assert().Equal("Test", T("Test"), "Expected translation to EN")
	t.Assert().Equal(EN, GetCurrentLanguage(), "Expected current locale to be EN")

	SetCurrentLocale(RU.String())
	t.Assert().Equal("Тест", T("Test"), "Expected translation to RU")
	t.Assert().Equal(RU, GetCurrentLanguage(), "Expected current locale to be RU")
}

func TestI18nSuite(t *testing.T) {
	suite.RunSuite(t, new(I18nSuite))
}
