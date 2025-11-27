package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Krzysztofz01/fns/printer"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the FNS version",
	Long:  "Print the semantic versioning of the FNS tool",
	Run: func(cmd *cobra.Command, args []string) {
		printer.GetPrinter().Printf("%s\n", Version)
	},
}
