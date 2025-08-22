package i18n

import (
	"os"
	"path"
	"strings"

	"testing"

	"github.com/joho/godotenv"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type HelpersSuite struct {
	suite.Suite
}

func (s *HelpersSuite) BeforeAll(t provider.T) {
	_ = godotenv.Load("../.env")
}

func (s *HelpersSuite) BeforeEach(t provider.T) {
	t.Epic("i18n")
	t.Feature("Helpers")
	t.Tags("i18n", "env")
	t.Severity(allure.NORMAL)
	t.Owner("github.com/11ALX11")
	t.Link(allure.LinkLink("PhraseApp-Blog i18n github", "https://github.com/PhraseApp-Blog/go-internationalization/tree/master/pkg/i18n"))
}

func (s *HelpersSuite) TestGetLocalePath(t provider.T) {
	t.Title("Test getLocalePath()")
	t.Description("Check if getLocalePath() returns correct path")

	expectedPath := path.Join(os.Getenv("APPROOTDIR"), os.Getenv("LOCALESDIR"))
	localePath, err := getLocalePath()

	t.Assert().NoError(err, "Expected no error from getLocalePath()")

	t.Assert().True(strings.Contains(localePath, expectedPath),
		"Expected path %s to contain %s, but it does not", localePath, expectedPath)
}

func (s *HelpersSuite) TestIsLocaleDirExists(t provider.T) {
	t.Title("Check locale dir")
	t.Description("Check if locale dir exists")

	localePath, err := getLocalePath()
	t.Assert().NoError(err, "Expected no error from getLocalePath()")

	_, statErr := os.Stat(localePath)
	t.Assert().False(os.IsNotExist(statErr),
		"Expected directory %s to exist, but it does not", localePath)
}

func (s *HelpersSuite) TestGetPwdDirPath(t provider.T) {
	t.Title("Test getPwdDirPath()")
	t.Description("Check if getPwdDirPath() returns correct root path")

	expectedPath := os.Getenv("APPROOTDIR")
	t.WithParameters(allure.NewParameter("APPROOTDIR", expectedPath))

	rootPath, err := getPwdDirPath()
	t.Assert().NoError(err, "Expected no error from getPwdDirPath()")

	t.Assert().True(strings.HasSuffix(rootPath, expectedPath),
		"Expected path %s to end with %s, but it does not", rootPath, expectedPath)
}

func (s *HelpersSuite) TestGetPwdDirPathWithEmptyEnv(t provider.T) {
	t.Title("Test getPwdDirPath()")
	t.Description("Check if getPwdDirPath() returns error when APPROOTDIR is empty")

	t.Setenv("APPROOTDIR", "")
	t.WithParameters(allure.NewParameter("APPROOTDIR", ""))

	path, err := getPwdDirPath()
	t.Assert().Error(err, "Expected error when APPROOTDIR is empty")
	t.Assert().Equal("", path, "Expected path to be empty")
}

func TestHelpersSuite(t *testing.T) {
	suite.RunSuite(t, new(HelpersSuite))
}
