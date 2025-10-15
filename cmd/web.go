package cmd

import (
	"fmt"

	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Use a web-based interface",
	// Long:  `Use a web-based interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(i18n.T("web not supported yet."))
	},
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
