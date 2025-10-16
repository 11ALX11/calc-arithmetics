package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/11ALX11/calc-arithmetics/i18n"
)

func runApp() {

	var content, keyPath, dataFileInArchive string
	var err error

	inputFile := prompt(
		i18n.T("Path to file to read input data from:\n"),
	)

	unzip := promptB(
		i18n.T("Is input file an archive? (y/N): "),
	)

	if !unzip {

		content, err = app.ReadFile(inputFile)

		if err != nil {
			log.Fatalf("Failed to read a file: %s; error: %s", inputFile, err)
			return
		}
	} else {

		dataFileInArchive = prompt(
			i18n.T("Name of the file inside the ZIP to read (data.txt):\n"),
		)
		if dataFileInArchive == "" {
			dataFileInArchive = app.DataFileInArchive
		}

		content, err = app.ReadZipFile(inputFile, dataFileInArchive)

		if err != nil {
			log.Fatalf("Failed to read an archive: %s; error: %s", inputFile, err)
			return
		}
	}

	decrypt := promptB(
		i18n.T("Is input file encoded? (y/N): "),
	)

	if decrypt {
		keyPath = prompt(
			i18n.T("Path to file with encryption key to decode with:\n"),
		)

		content, err = app.DecryptFileKey(content, keyPath)

		if err != nil {
			log.Fatalf("Failed to decipher, error: %s", err)
			return
		}
	}

	useEvalLib := promptB(
		i18n.T("Use evaluation library? (y/N): "),
	)

	evalFunction := app.Eval
	if useEvalLib {
		evalFunction = app.EvalLib
	}

	useFilterRegex := promptB(
		i18n.T("Use regex for filtering arithmetic expressions? (y/N): "),
	)

	replaceFunction := app.ReplaceMathExpressions
	if useFilterRegex {
		replaceFunction = app.ReplaceMathExpressionsRegex
	}

	sResult := replaceFunction(content, evalFunction)

	encrypt := promptB(
		i18n.T("Do you wish to encrypt result? (y/N): "),
	)

	if encrypt {
		if keyPath == "" {
			keyPath = prompt(
				i18n.T("Path to file with encryption key to decode with:\n"),
			)
		}
		sResult, err = app.EncryptFileKey(sResult, keyPath)

		if err != nil {
			log.Fatalf("Failed to encode, error: %s", err)
			return
		}
	}

	outputFile := prompt(
		i18n.T("Path to file to write result to:\n"),
	)

	archive := promptB(
		i18n.T("Do you wish to archive output file? (y/N): "),
	)

	if archive {

		if dataFileInArchive == "" {
			dataFileInArchive = prompt(
				i18n.T("Name of the file inside the ZIP to write (data.txt):\n"),
			)
			if dataFileInArchive == "" {
				dataFileInArchive = app.DataFileInArchive
			}
		}

		err = app.WriteZipFile(outputFile, sResult, dataFileInArchive)
	} else {

		err = app.WriteFile(outputFile, sResult)
	}

	if err != nil {
		log.Fatalf("Failed to write a file: %s; error: %s", outputFile, err)
		return
	}

	outputToConsole := promptB(
		i18n.T("Also print the results to the console? (y/N): "),
	)

	if outputToConsole {
		fmt.Println(sResult)
	}
}

func prompt(q string) string {
	fmt.Print(q)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to prompt: %s; error: %s", q, err)
	}
	return strings.TrimSpace(str)
}

func promptB(q string) bool {
	str := prompt(q)
	return strings.Contains(strings.ToLower(str), "y")
}
