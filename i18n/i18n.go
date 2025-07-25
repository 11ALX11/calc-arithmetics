/*
 * source: https://github.com/PhraseApp-Blog/go-internationalization/tree/master/pkg/i18n
 */
package i18n

import (
	"os"
	"strings"

	"github.com/leonelquinteros/gotext"
)

var (
	defaultDomain = "default"
)

func Init() error {
	localePath, err := getLocalePath()
	if err != nil {
		return err
	}
	languageCode := getLanguageCode()
	fullLocale := NewLanguageFromString(languageCode).String()
	gotext.Configure(localePath, fullLocale, defaultDomain)
	setupLocales(localePath)
	// fmt.Println("languageCode:", localePath)

	return nil
}

// GetLanguageCode returns the language code from environment variables LANGUAGE, LC_ALL, or LC_MESSAGES,
// in that order of priority. It returns an empty string if none of the variables are set.
func getLanguageCode() string {
	// Check LANGUAGE environment variable
	if lc := os.Getenv("LANGUAGE"); lc != "" {
		return lc
	}
	// Check LC_ALL environment variable
	if lc := os.Getenv("LC_ALL"); lc != "" {
		return lc
	}
	// Check LC_MESSAGES environment variable
	if lc := os.Getenv("LC_MESSAGES"); lc != "" {
		return lc
	}
	// No language code found in environment variables
	return os.Getenv("LANG")
}

func setupLocales(localePath string) error {
	// Get a list of all directories in the locale path
	localeDirs, err := os.ReadDir(localePath)
	if err != nil {
		return err
	}

	// Iterate over each directory and add it as a supported language
	for _, dir := range localeDirs {
		if dir.IsDir() {
			langCode := LanguageCode(dir.Name())
			lang := gotext.NewLocale(localePath, langCode.String())
			lang.AddDomain(defaultDomain)
			langMap[langCode] = lang
		}
	}

	return nil
}

// GetSupportedLanguages returns the list of supported language codes.
func GetSupportedLanguages() []LanguageCode {
	// Initialize an empty slice to store language codes
	var languages []LanguageCode

	// Iterate over the keys of langMap and collect them into the slice
	for lang := range langMap {
		languages = append(languages, lang)
	}

	return languages
}

func NewLanguageFromString(code string) LanguageCode {
	code = strings.ToLower(code)
	if strings.Contains(code, "ru") {
		return RU
	}

	return EN
}

func T(s string, args ...interface{}) string {
	return gotext.Get(s, args...)
}

// *** Not used, Not tested ***
//
//	func TN(s string, p string, n int, args ...interface{}) string {
//		return gotext.GetN(s, p, n, n)
//	}

func GetCurrentLanguage() LanguageCode {
	// Get the current language code from the gotext package
	lang := gotext.GetLanguage()

	// Convert the language code to LanguageCode enum
	return LanguageCode(lang)
}

// *** Not used, Not tested ***
// func DetectPreferredLocale(r *http.Request) string {
// 	// Check if lang parameter is provided in the URL
// 	langParam := LanguageCode(r.URL.Query().Get("lang"))
// 	if langParam != "" {
// 		// Check if the provided lang parameter is supported
// 		for _, supportedLang := range GetSupportedLanguages() {
// 			if langParam == supportedLang {
// 				return langParam.String()
// 			}
// 		}
// 	}

// 	// Get Accept-Language header value
// 	acceptLanguage := r.Header.Get("Accept-Language")

// 	// Parse Accept-Language header
// 	prefs, _, err := language.ParseAcceptLanguage(acceptLanguage)
// 	if err != nil {
// 		// Default to English if parsing fails
// 		return "en_US"
// 	}

// 	// Convert supported language codes to language.Tags
// 	var supportedTags []language.Tag
// 	for _, code := range GetSupportedLanguages() {
// 		tag := language.Make(code.String())
// 		supportedTags = append(supportedTags, tag)
// 	}

// 	// Find the best match between supported languages and client preferences
// 	match := language.NewMatcher(supportedTags)
// 	_, index, _ := match.Match(prefs...)

// 	// Get the best match language
// 	locale := GetSupportedLanguages()[index]

// 	return locale.String()
// }

// SetCurrentLocale sets the current locale based on the language code
func SetCurrentLocale(lang string) {
	// If the language parameter is provided, set the current locale
	if lang != "" {
		// Get the preferred locale based on the language code
		locale := (lang)

		// Set the current locale
		gotext.SetLanguage(locale)
	}
}

// *** Not used, Not tested ***
// func FormatLocalizedDate(t time.Time, lang string) string {
// 	// Read date formats from JSON file
// 	dateFormats, err := readDateFormatsFromFile("date_formats.json")
// 	if err != nil {
// 		// Log or handle error
// 		// Fallback to default format if unable to read formats
// 		return t.Format("02/01/2006 15:04:05")
// 	}

// 	// Get the appropriate date format for the given language
// 	format, ok := dateFormats[lang]
// 	if !ok {
// 		// If the language is not recognized, use a default format
// 		return t.Format("02/01/2006 15:04:05")
// 	}

// 	// Load the appropriate time location based on the language tag
// 	loc, err := time.LoadLocation(lang)
// 	if err != nil {
// 		// Log or handle error
// 		// Fallback to default location if an error occurs
// 		loc = time.UTC
// 	}
// 	fmt.Println(format)
// 	// Format the time using the specified format and location
// 	return t.In(loc).Format(format)
// }

// *** Not used, Not tested ***
// // readDateFormatsFromFile reads date formats from a JSON file.
// func readDateFormatsFromFile(filename string) (map[string]string, error) {
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var dateFormats map[string]string
// 	if err := json.Unmarshal(data, &dateFormats); err != nil {
// 		return nil, err
// 	}

// 	return dateFormats, nil
// }

// *** Not used, Not tested ***
// FormatNumber formats the given number according to the specified language locale.
// func FormatNumber(number int64, lang language.Tag) string {

// 	// Create a new message printer with the specified language
// 	p := message.NewPrinter(lang)
// 	// Format the number with grouping separators according to the user's preferred language
// 	return p.Sprintf("%d", number)
// }
