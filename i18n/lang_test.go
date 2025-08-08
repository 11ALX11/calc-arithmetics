package i18n

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type LanguageSuite struct {
	suite.Suite
}

func (s *LanguageSuite) BeforeEach(t provider.T) {
	t.Epic("i18n")
	t.Feature("Language")
	t.Tags("i18n")
}

func (s *LanguageSuite) TestLanguageT(t provider.T) {
	t.Description("Test Language.T() for EN and RU returns expected strings")
	t.Tag("env")

	_ = godotenv.Load("../.env")
	Init()

	expectedEN := "Test"
	gotEN := EN.T("Test")
	t.Assert().Equal(gotEN, expectedEN, "Expected translation to EN to be %s", expectedEN)

	expectedRU := "Тест"
	gotRU := RU.T("Test")
	t.Assert().Equal(gotRU, expectedRU, "Expected translation to RU to be %s", expectedRU)
}

func (s *LanguageSuite) TestString(t provider.T) {
	t.Description("Test String() for LanguageCode EN and RU")

	t.Assert().Equal(EN.String(), "en_US", "Expected EN.String() to be en_US")
	t.Assert().Equal(RU.String(), "ru_RU", "Expected RU.String() to be ru_RU")
}

func TestLanguageSuite(t *testing.T) {
	suite.RunSuite(t, new(LanguageSuite))
}
