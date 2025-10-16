package console

import (
	"fmt"

	"github.com/11ALX11/calc-arithmetics/flags"
	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	// flag: useOop
	if flags.UseOop {
		//NewConsoleRunner(args).Run()
		fmt.Println(i18n.T("console not supported yet."))
	} else {
		runApp()
	}
}
