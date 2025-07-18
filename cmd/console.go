package cmd

import (
	"fmt"

	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: i18n.T("Use a console-based interface"),
	// Long:  `Use a console-based interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(i18n.T("console not supported yet."))
	},
}

func init() {
	rootCmd.AddCommand(consoleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consoleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consoleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
