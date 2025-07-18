package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// guiCmd represents the gui command
var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Use a graphic user interface",
	// Long: `Use an user interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gui not supported yet.")
	},
}

func init() {
	rootCmd.AddCommand(guiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// guiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
