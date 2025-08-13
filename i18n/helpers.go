/*
 * source: https://github.com/PhraseApp-Blog/go-internationalization/tree/master/pkg/i18n
 */
package i18n

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func getLocalePath() (string, error) {
	rootPath, err := getPwdDirPath()
	if err != nil {
		return "", err
	}
	return path.Join(rootPath, os.Getenv("LOCALESDIR")), nil
}

func getPwdDirPath() (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	expectedSuffix := os.Getenv("APPROOTDIR")

	// Assume that we can reach app root dir by climbing FS heirarchy
	pattern := fmt.Sprintf(`^(.*)(%s)`, regexp.QuoteMeta(expectedSuffix))

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(rootPath)

	if len(matches) == 0 || matches[0] == "" || !strings.HasSuffix(matches[0], expectedSuffix) {
		return "", fmt.Errorf("path is empty or does not contain %s", expectedSuffix)
	}

	return matches[0], nil
}
