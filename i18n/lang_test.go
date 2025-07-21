package i18n

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestLanguageT(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file")
		return
	}

	Init()

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

func TestString(t *testing.T) {
	if EN.String() != "en_US" {
		t.Errorf("Expected LanguageCode EN to have string representation en_US, but it does not")
	}

	if RU.String() != "ru_RU" {
		t.Errorf("Expected LanguageCode RU to have string representation ru_RU, but it does not")
	}
}
