package cmd

import (
	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/11ALX11/calc-arithmetics/cli"
	"github.com/11ALX11/calc-arithmetics/flags"
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

	cliCmd.Flags().BoolVarP(&flags.OutputToConsole, "outputToConsole", "o", false, "Also print the results to the console")

	cliCmd.PersistentFlags().BoolVarP(&flags.UseEvalLib, "evalLib", "e", false, "Use an evaluation library expr-lang.")
	cliCmd.PersistentFlags().BoolVarP(&flags.UseFilterRegex, "filterRegex", "f", false, "Use regex for filtering arithmetic expressions from file.")

	cliCmd.PersistentFlags().BoolVarP(&flags.Unzip, "unzip", "u", false, "Read input file from a ZIP archive")
	cliCmd.PersistentFlags().BoolVarP(&flags.Archive, "archive", "a", false, "Write output file as a ZIP archive")
	cliCmd.PersistentFlags().StringVarP(
		&flags.DataFileInArchive,
		"dataFileInArchive",
		"d",
		app.DataFileInArchive,
		"Name of the file inside the ZIP to read (with --unzip) or write (with --archive). Can be used with either or both --unzip and --archive.\nWhen reading, if the specified file is not present, the first file in the archive is used.",
	)

	cliCmd.PersistentFlags().BoolVar(&flags.Decrypt, "decrypt", false, "Decrypt input file. Use with --keyPath")
	cliCmd.PersistentFlags().BoolVar(&flags.Encrypt, "encrypt", false, "Encrypt output file. Use with --keyPath")
	cliCmd.PersistentFlags().StringVar(&flags.KeyPath, "keyPath", "", "Path to a file containing the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.\nRequired when --encrypt or --decrypt is set")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
