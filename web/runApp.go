package web

import (
	"fmt"

	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/11ALX11/calc-arithmetics/flags"
)

func replaceMathExpressions(expr string) (string, error) {

	content := expr
	var err error

	// flag: keyPath. Check if set
	if (flags.Decrypt || flags.Encrypt) && flags.KeyPath == "" {
		return "", fmt.Errorf("keyPath is not set")
	}

	// flag: unzip
	if flags.Unzip {

		content, err = app.ReadZipData(content, flags.DataFileInArchive)

		if err != nil {
			return "", fmt.Errorf("failed to read an archive: %s; error: %s", expr, err)
		}
	}

	// flag: decrypt
	if flags.Decrypt {
		content, err = app.DecryptFileKey(content, flags.KeyPath)

		if err != nil {
			return "", fmt.Errorf("failed to decipher, error: %s", err)
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
			return "", fmt.Errorf("failed to encode, error: %s", err)
		}
	}

	// flag: archive
	if flags.Archive {
		archived, err := app.GetZipData(sResult, flags.DataFileInArchive)

		if err != nil {
			return "", fmt.Errorf("failed to archive a string: %s; error: %s", sResult, err)
		}

		return archived, nil
	}

	return sResult, nil
}
