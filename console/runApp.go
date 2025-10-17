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

const (
	promptInputFile         = "Path to file to read input data from:\n"
	promptIsInputArchive    = "Is input file an archive? (y/N): "
	promptDataFileInArchive = "Name of the file inside the ZIP to read (data.txt):\n"
	promptIsInputEncoded    = "Is input file encoded? (y/N): "
	promptKeyPath           = "Path to file with encryption key to:\n"
	promptUseEvalLib        = "Use evaluation library? (y/N): "
	promptUseFilterRegex    = "Use regex for filtering arithmetic expressions? (y/N): "
	promptEncryptResult     = "Do you wish to encrypt result? (y/N): "
	promptOutputFile        = "Path to file to write result to:\n"
	promptArchiveOutput     = "Do you wish to archive output file? (y/N): "
	promptOutputToConsole   = "Also print the results to the console? (y/N): "
)

func runApp() {

	var content, keyPath, dataFileInArchive string
	var err error

	inputFile := prompt(i18n.T(promptInputFile))

	unzip := promptB(i18n.T(promptIsInputArchive))

	if !unzip {
		content, err = app.ReadFile(inputFile)

		if err != nil {
			log.Fatalf("Failed to read a file: %s; error: %s", inputFile, err)
			return
		}
	} else {
		dataFileInArchive = prompt(i18n.T(promptDataFileInArchive))
		if dataFileInArchive == "" {
			dataFileInArchive = app.DataFileInArchive
		}

		content, err = app.ReadZipFile(inputFile, dataFileInArchive)

		if err != nil {
			log.Fatalf("Failed to read an archive: %s; error: %s", inputFile, err)
			return
		}
	}

	decrypt := promptB(i18n.T(promptIsInputEncoded))

	if decrypt {
		keyPath = prompt(i18n.T(promptKeyPath))

		content, err = app.DecryptFileKey(content, keyPath)

		if err != nil {
			log.Fatalf("Failed to decipher, error: %s", err)
			return
		}
	}

	useEvalLib := promptB(i18n.T(promptUseEvalLib))

	evalFunction := app.Eval
	if useEvalLib {
		evalFunction = app.EvalLib
	}

	useFilterRegex := promptB(i18n.T(promptUseFilterRegex))

	replaceFunction := app.ReplaceMathExpressions
	if useFilterRegex {
		replaceFunction = app.ReplaceMathExpressionsRegex
	}

	sResult := replaceFunction(content, evalFunction)

	encrypt := promptB(i18n.T(promptEncryptResult))

	if encrypt {
		if keyPath == "" {
			keyPath = prompt(i18n.T(promptKeyPath))
		}
		sResult, err = app.EncryptFileKey(sResult, keyPath)

		if err != nil {
			log.Fatalf("Failed to encode, error: %s", err)
			return
		}
	}

	outputFile := prompt(i18n.T(promptOutputFile))

	archive := promptB(i18n.T(promptArchiveOutput))

	if archive {
		if dataFileInArchive == "" {
			dataFileInArchive = prompt(i18n.T(promptDataFileInArchive))
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

	outputToConsole := promptB(i18n.T(promptOutputToConsole))

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
