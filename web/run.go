package web

import (
	"fmt"

	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println(i18n.T("web not supported yet."))
}
