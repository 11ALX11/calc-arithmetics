package console

import (
	"github.com/11ALX11/calc-arithmetics/flags"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	// flag: useOop
	if flags.UseOop {
		NewConsoleRunner().Run()
	} else {
		runApp()
	}
}
