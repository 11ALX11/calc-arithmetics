package cmd

import (
	"os"

	"github.com/11ALX11/calc-arithmetics/app"
	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calc-arithmetics",
	Short: i18n.T("Find all arithmetic operations, calculate and replace."),
	Long:  i18n.T(`Find all arithmetic operations in the input file, calculate and replace with the results in the output file.`),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// flag: forceTranslation
		if forceTranslation != "" {
			i18n.SetCurrentLocale(forceTranslation)
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

var (
	forceTranslation string
	outputToConsole  bool
	verboseOutput    bool

	useEvalLib     bool
	useFilterRegex bool

	unzip             bool
	archive           bool
	dataFileInArchive string

	decrypt bool
	encrypt bool
	keyPath string
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calc-arithmetics.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&forceTranslation, "forceTranslation", "t", "", "Forces translation of the app to the preferred locale. Options: \"en_US\" | \"ru_RU\"")
	rootCmd.PersistentFlags().BoolVarP(&outputToConsole, "outputToConsole", "o", false, "Also print the results to the console")
	rootCmd.PersistentFlags().BoolVarP(&verboseOutput, "verbose", "v", false, "Enable verbose (info-level) logging")

	rootCmd.PersistentFlags().BoolVarP(&useEvalLib, "evalLib", "e", false, "Use an evaluation library expr-lang.")
	rootCmd.PersistentFlags().BoolVarP(&useFilterRegex, "filterRegex", "f", false, "Use regex for filtering arithmetic expressions from file.")

	rootCmd.PersistentFlags().BoolVarP(&unzip, "unzip", "u", false, "Read input file from a ZIP archive")
	rootCmd.PersistentFlags().BoolVarP(&archive, "archive", "a", false, "Write output file as a ZIP archive")
	rootCmd.PersistentFlags().StringVarP(
		&dataFileInArchive,
		"dataFileInArchive",
		"d",
		app.DataFileInArchive,
		"Name of the file inside the ZIP to read (with --unzip) or write (with --archive). Can be used with either or both --unzip and --archive",
	)

	rootCmd.PersistentFlags().BoolVar(&decrypt, "decrypt", false, "Decrypt input file. Use with --keyPath")
	rootCmd.PersistentFlags().BoolVar(&encrypt, "encrypt", false, "Encrypt output file. Use with --keyPath")
	rootCmd.PersistentFlags().StringVar(&keyPath, "keyPath", "", "Path to a file containing the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.\nRequired when --encrypt or --decrypt is set")
}
