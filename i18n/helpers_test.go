package i18n

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetLocalePath(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file")
		return
	}

	expectedPath := path.Join(os.Getenv("APPROOTDIR"), os.Getenv("LOCALESDIR"))

	localePath, err := getLocalePath()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.Contains(localePath, expectedPath) {
		t.Errorf("Expected path %s to contain %s, but it does not", localePath, expectedPath)
	}
}

func TestIsLocaleDirExists(t *testing.T) {
	localePath, err := getLocalePath()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if _, err := os.Stat(localePath); os.IsNotExist(err) {
		t.Errorf("Expected directory %s to exist, but it does not", localePath)
	}
}

func TestGetPwdDirPath(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Fatalf("Error loading .env file")
		return
	}

	expectedPath := os.Getenv("APPROOTDIR")

	rootPath, err := getPwdDirPath()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !strings.HasSuffix(rootPath, expectedPath) {
		t.Errorf("Expected path %s to end with %s, but it does not", rootPath, expectedPath)
	}
}

func TestGetPwdDirPathWithEmptyEnv(t *testing.T) {
	t.Setenv("APPROOTDIR", "")

	path, err := getPwdDirPath()
	if err == nil {
		t.Errorf("Expected error, got nil; resulting path: %s", path)
	}
}
