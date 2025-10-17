package console

import (
	"fmt"
	"log"

	"github.com/11ALX11/calc-arithmetics/app"
	app_oop "github.com/11ALX11/calc-arithmetics/app-oop"
	"github.com/11ALX11/calc-arithmetics/i18n"
)

type ConsoleRunner struct{}

func NewConsoleRunner() *ConsoleRunner {
	return &ConsoleRunner{}
}

func (c ConsoleRunner) Run() {

	var keyPath, dataFileInArchive string

	inputFile := c.prompt(i18n.T(promptInputFile))

	reader := app_oop.NewReadin()

	unzip := c.promptB(i18n.T(promptIsInputArchive))
	if unzip {

		dataFileInArchive = prompt(i18n.T(promptDataFileInArchive))
		if dataFileInArchive == "" {
			dataFileInArchive = app.DataFileInArchive
		}

		reader = app_oop.NewUnzip(reader, dataFileInArchive)
	}

	decrypt := c.promptB(i18n.T(promptIsInputEncoded))
	if decrypt {

		keyPath = prompt(i18n.T(promptKeyPath))

		reader = app_oop.NewDecrypt(reader, keyPath)
	}

	content, err := reader.
		ReadFile(inputFile).
		GetContentError()

	if err != nil {
		log.Fatalf("Failed to read file: %s; error: %s", inputFile, err)
		return
	}

	useEvalLib := c.promptB(i18n.T(promptUseEvalLib))
	useFilterRegex := c.promptB(i18n.T(promptUseFilterRegex))

	sResult := app_oop.
		NewFilterFactory(useFilterRegex).
		GetFilterImplementation(app_oop.
			NewEvalFactory(useEvalLib).
			GetEvalImplementation(),
		).
		ReplaceMathExpressions(content)

	writer := app_oop.NewWriteout()

	archive := promptB(i18n.T(promptArchiveOutput))
	if archive {

		if dataFileInArchive == "" {
			dataFileInArchive = prompt(i18n.T(promptDataFileInArchive))
			if dataFileInArchive == "" {
				dataFileInArchive = app.DataFileInArchive
			}
		}

		writer = app_oop.NewArchive(writer, dataFileInArchive)
	}

	encrypt := c.promptB(i18n.T(promptEncryptResult))
	if encrypt {

		if keyPath == "" {
			keyPath = prompt(i18n.T(promptKeyPath))
		}

		writer = app_oop.NewEncrypt(writer, keyPath)
	}

	outputFile := c.prompt(i18n.T(promptOutputFile))
	writer.WriteFile(outputFile, sResult)

	if writer.GetError() != nil {
		log.Fatalf("Failed to write a file: %s; error: %s", outputFile, writer.GetError())
		return
	}

	outputToConsole := c.promptB(i18n.T(promptOutputToConsole))
	if outputToConsole {
		fmt.Println(sResult)
	}
}

func (c ConsoleRunner) prompt(q string) string {
	return prompt(q)
}

func (c ConsoleRunner) promptB(q string) bool {
	return promptB(q)
}
