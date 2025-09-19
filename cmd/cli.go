package cmd

import (
	"fmt"
	"log"

	"github.com/11ALX11/calc-arithmetics/app"
	app_oop "github.com/11ALX11/calc-arithmetics/app-oop"
	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli INPUT_FILE OUTPUT_FILE",
	Short: i18n.T("Use a command-line interface"),
	// Long:  `Use command-line interface.`,
	Args: cobra.ExactArgs(2),
	Run:  run,
}

func init() {
	rootCmd.AddCommand(cliCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run(cmd *cobra.Command, args []string) {
	// flag: useOop
	if useOop {
		runAppOop(args)
	} else {
		runApp(args)
	}
}

func runApp(args []string) {

	var content string
	var err error

	// flag: keyPath. Check if set
	if (decrypt || encrypt) && keyPath == "" {
		log.Fatalf("keyPath is not set.")
		return
	}

	// Read normally
	if !unzip {

		content, err = app.ReadFile(args[0])

		if err != nil {
			log.Fatalf("Failed to read a file: %s; error: %s", args[0], err)
			return
		}
	} else {
		// flag: unzip

		content, err = app.ReadZipFile(args[0], dataFileInArchive)

		if err != nil {
			log.Fatalf("Failed to read an archive: %s; error: %s", args[0], err)
			return
		}
	}

	// flag: decrypt
	if decrypt {
		content, err = app.DecryptFileKey(content, keyPath)

		if err != nil {
			log.Fatalf("Failed to decipher, error: %s", err)
			return
		}
	}

	// flag: useEvalLib
	evalFunction := app.Eval
	if useEvalLib {
		evalFunction = app.EvalLib
	}

	// flag: useFilterRegex
	replaceFunction := app.ReplaceMathExpressions
	if useFilterRegex {
		replaceFunction = app.ReplaceMathExpressionsRegex
	}

	sResult := replaceFunction(content, evalFunction)

	// flag: encrypt
	if encrypt {
		sResult, err = app.EncryptFileKey(sResult, keyPath)

		if err != nil {
			log.Fatalf("Failed to encode, error: %s", err)
			return
		}
	}

	// flag: outputToConsole
	if outputToConsole {
		fmt.Println(sResult)
	}

	// flag: archive
	if archive {
		err = app.WriteZipFile(args[1], sResult, dataFileInArchive)
	} else {
		// Write normally

		err = app.WriteFile(args[1], sResult)
	}

	if err != nil {
		log.Fatalf("Failed to write a file: %s; error: %s", args[1], err)
		return
	}
}

func runAppOop(args []string) {

	// flag: keyPath. Check if set
	if (decrypt || encrypt) && keyPath == "" {
		log.Fatalf("keyPath is not set.")
		return
	}

	// flag: unzip, dataFileInArchive
	reader := app_oop.
		NewReaderFactory(unzip).
		GetReaderImplementation().
		SetDataInputFile(dataFileInArchive).
		ReadFile(args[0])

	// flag: decrypt
	if decrypt {
		reader.Accept(app_oop.
			NewDecrypt().
			SetKeyPath(keyPath),
		)
	}

	content, err := reader.GetContentError()

	if err != nil {
		log.Fatalf("Failed to read file: %s; error: %s", args[0], err)
		return
	}

	// flag: useEvalLib
	// flag: useFilterRegex
	sResult := app_oop.
		NewFilterFactory(useFilterRegex).
		GetFilterImplementation().
		SetEvalFuncWithEvaluator(app_oop.
			NewEvalFactory(useEvalLib).
			GetEvalImplementation(),
		).
		ReplaceMathExpressions(content)

	// flag: encrypt
	if encrypt {
		sResult, err = app.EncryptFileKey(sResult, keyPath)

		if err != nil {
			log.Fatalf("Failed to encode, error: %s", err)
			return
		}
	}

	// flag: outputToConsole
	if outputToConsole {
		fmt.Println(sResult)
	}

	// flag: archive
	if archive {
		err = app.WriteZipFile(args[1], sResult, dataFileInArchive)
	} else {
		// Write normally

		err = app.WriteFile(args[1], sResult)
	}

	if err != nil {
		log.Fatalf("Failed to write a file: %s; error: %s", args[1], err)
		return
	}
}
