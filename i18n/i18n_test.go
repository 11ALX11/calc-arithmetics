package i18n

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestInit(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file")
		return
	}

	if err := Init(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len := len(langMap); len < 2 {
		t.Errorf("Expected to have at least 2 locales in a map, got %d", len)
	}

	expectedStr := "Test"
	translationStr := EN.T("Test")
	if translationStr != expectedStr {
		t.Errorf("Expected translation to EN \"%s\" to match \"%s\", but it does not", translationStr, expectedStr)
	}

	expectedStr = "Тест"
	translationStr = RU.T("Test")
	if translationStr != expectedStr {
		t.Errorf("Expected translation to RU \"%s\" to match \"%s\", but it does not", translationStr, expectedStr)
	}
}

func TestGetLanguageCodeLANGUAGE(t *testing.T) {
	t.Setenv("LANGUAGE", "ru_RU")

	lang := getLanguageCode()
	if lang != "ru_RU" {
		t.Errorf("Expected to get language code ru_RU, got %s", lang)
	}
}

func TestGetLanguageCodeLC_ALL(t *testing.T) {
	t.Setenv("LC_ALL", "ru_RU")

	lang := getLanguageCode()
	if lang != "ru_RU" {
		t.Errorf("Expected to get language code ru_RU, got %s", lang)
	}
}

func TestGetLanguageCodeLC_MESSAGES(t *testing.T) {
	t.Setenv("LC_MESSAGES", "ru_RU")

	lang := getLanguageCode()
	if lang != "ru_RU" {
		t.Errorf("Expected to get language code ru_RU, got %s", lang)
	}
}

func TestGetLanguageCodeLANG(t *testing.T) {
	t.Setenv("LANG", "ru_RU")

	lang := getLanguageCode()
	if lang != "ru_RU" {
		t.Errorf("Expected to get language code ru_RU, got %s", lang)
	}
}

func TestGetSupportedLanguages(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file")
		return
	}

	Init()

	if len := len(GetSupportedLanguages()); len < 2 {
		t.Errorf("Expected at least 2 supported languages, got %d", len)
	}
}

func TestNewLanguageFromString(t *testing.T) {
	if lang := NewLanguageFromString("ru_RU"); lang != RU {
		t.Errorf("Expected for ru_RU to match RU enum, got %s", lang)
	}

	if lang := NewLanguageFromString("ru"); lang != RU {
		t.Errorf("Expected for ru to match RU enum, got %s", lang)
	}

	if lang := NewLanguageFromString("en_US"); lang != EN {
		t.Errorf("Expected for en_US to match EN enum, got %s", lang)
	}
}

func TestGetCurrentLanguage(t *testing.T) {
	t.Setenv("LANGUAGE", "ru_RU")
	Init()

	lang := GetCurrentLanguage()
	if lang != RU {
		t.Errorf("Expected to get language (enum) RU, got %s", lang)
	}
}

func TestT(t *testing.T) {
	t.Setenv("LANGUAGE", "ru_RU")

	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file")
		return
	}

	if err := Init(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedStr := "Тест"
	translationStr := T("Test")
	if translationStr != expectedStr {
		t.Errorf("Expected translation to RU \"%s\" to match \"%s\", but it does not", translationStr, expectedStr)
	}
}

func TestSetCurrentLocale(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file")
		return
	}

	if err := Init(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Setenv("LANGUAGE", "ru_RU")

	SetCurrentLocale(EN.String())
	if lang := GetCurrentLanguage(); lang != EN {
		t.Errorf("Expected switch to EN locale, current locale: %s", lang)
	}

	SetCurrentLocale(RU.String())
	if lang := GetCurrentLanguage(); lang != RU {
		t.Errorf("Expected switch to RU locale, current locale: %s", lang)
	}
}
