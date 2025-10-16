package cmd

import (
	"os"

	"github.com/11ALX11/calc-arithmetics/flags"
	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calc-arithmetics",
	Short: "Find all arithmetic operations, calculate and replace.",
	Long:  `Find all arithmetic operations in the input file, calculate and replace with the results in the output file.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// flag: forceTranslation
		if flags.ForceTranslation != "" {
			i18n.SetCurrentLocale(flags.ForceTranslation)
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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calc-arithmetics.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&flags.ForceTranslation, "forceTranslation", "t", "", "Forces translation of the app to the preferred locale. Options: \"en_US\" | \"ru_RU\"")
	rootCmd.PersistentFlags().BoolVarP(&flags.VerboseOutput, "verbose", "v", false, "Enable verbose (info-level) logging")
	rootCmd.PersistentFlags().BoolVar(&flags.UseOop, "oop", false, "Use the object‑oriented engine instead of the default engine")
}
