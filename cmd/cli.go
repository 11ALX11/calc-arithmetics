package cmd

import (
	"fmt"

	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: i18n.T("Use a command-line interface"),
	// Long:  `Use command-line interface.`,
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {

		content, err := app.ReadFile(args[0])
		if err != nil {
			fmt.Errorf("failed to read a file: %s; error: %s", args[0], err)
			return
		}

		err = app.WriteFile(args[1], app.ReplaceMathExpressions(content))
		if err != nil {
			fmt.Errorf("failed to write a file: %s; error: %s", args[1], err)
			return
		}
	},
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
