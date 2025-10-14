package cli

import (
	"fmt"
	"log"

	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/11ALX11/calc-arithmetics/flags"
)

func runApp(args []string) {

	var content string
	var err error

	// flag: keyPath. Check if set
	if (flags.Decrypt || flags.Encrypt) && flags.KeyPath == "" {
		log.Fatalf("keyPath is not set.")
		return
	}

	// Read normally
	if !flags.Unzip {

		content, err = app.ReadFile(args[0])

		if err != nil {
			log.Fatalf("Failed to read a file: %s; error: %s", args[0], err)
			return
		}
	} else {
		// flag: unzip

		content, err = app.ReadZipFile(args[0], flags.DataFileInArchive)

		if err != nil {
			log.Fatalf("Failed to read an archive: %s; error: %s", args[0], err)
			return
		}
	}

	// flag: decrypt
	if flags.Decrypt {
		content, err = app.DecryptFileKey(content, flags.KeyPath)

		if err != nil {
			log.Fatalf("Failed to decipher, error: %s", err)
			return
		}
	}

	// flag: useEvalLib
	evalFunction := app.Eval
	if flags.UseEvalLib {
		evalFunction = app.EvalLib
	}

	// flag: useFilterRegex
	replaceFunction := app.ReplaceMathExpressions
	if flags.UseFilterRegex {
		replaceFunction = app.ReplaceMathExpressionsRegex
	}

	sResult := replaceFunction(content, evalFunction)

	// flag: encrypt
	if flags.Encrypt {
		sResult, err = app.EncryptFileKey(sResult, flags.KeyPath)

		if err != nil {
			log.Fatalf("Failed to encode, error: %s", err)
			return
		}
	}

	// flag: outputToConsole
	if flags.OutputToConsole {
		fmt.Println(sResult)
	}

	// flag: archive
	if flags.Archive {
		err = app.WriteZipFile(args[1], sResult, flags.DataFileInArchive)
	} else {
		// Write normally

		err = app.WriteFile(args[1], sResult)
	}

	if err != nil {
		log.Fatalf("Failed to write a file: %s; error: %s", args[1], err)
		return
	}
}
