package cli

import (
	"fmt"
	"log"

	app_oop "github.com/11ALX11/calc-arithmetics/app-oop"
	"github.com/11ALX11/calc-arithmetics/flags"
)

func runAppOop(args []string) {

	// flag: keyPath. Check if set
	if (flags.Decrypt || flags.Encrypt) && flags.KeyPath == "" {
		log.Fatalf("keyPath is not set.")
		return
	}

	reader := app_oop.NewReadin()

	// flag: unzip, dataFileInArchive
	if flags.Unzip {
		reader = app_oop.NewUnzip(reader, flags.DataFileInArchive)
	}

	// flag: decrypt
	if flags.Decrypt {
		reader = app_oop.NewDecrypt(reader, flags.KeyPath)
	}

	content, err := reader.
		ReadFile(args[0]).
		GetContentError()

	if err != nil {
		log.Fatalf("Failed to read file: %s; error: %s", args[0], err)
		return
	}

	// flag: useEvalLib
	// flag: useFilterRegex
	sResult := app_oop.
		NewFilterFactory(flags.UseFilterRegex).
		GetFilterImplementation(app_oop.
			NewEvalFactory(flags.UseEvalLib).
			GetEvalImplementation(),
		).
		ReplaceMathExpressions(content)

	// flag: outputToConsole
	if flags.OutputToConsole {
		fmt.Println(sResult)
	}

	writer := app_oop.NewWriteout()

	// flag: archive
	if flags.Archive {
		writer = app_oop.NewArchive(writer, flags.DataFileInArchive)
	}

	// flag: encrypt
	if flags.Encrypt {
		writer = app_oop.NewEncrypt(writer, flags.KeyPath)
	}

	writer.WriteFile(args[1], sResult)

	if writer.GetError() != nil {
		log.Fatalf("Failed to write a file: %s; error: %s", args[1], err)
		return
	}
}
