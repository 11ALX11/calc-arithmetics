package cmd

import (
	"github.com/11ALX11/calc-arithmetics/cli"
	"github.com/spf13/cobra"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli INPUT_FILE OUTPUT_FILE",
	Short: "Use a command-line interface",
	// Long:  `Use command-line interface.`,
	Args: cobra.ExactArgs(2),
	Run:  cli.Run,
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
