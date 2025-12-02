package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/printer"
)

var (
	Default bool
)

func init() {
	configCmd.PersistentFlags().BoolVarP(&Default, "default", "d", false, "Print the default configuration instead of the current.")

	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the configuration of the FNS tool",
	Long:  "Print the JSON configuration of the FNS tool to the standard output",
	RunE: func(cmd *cobra.Command, args []string) error {
		var configuration *config.Configuration
		if Default {
			configuration = config.GetDefaultConfiguration()
		} else {
			configuration = config.GetConfiguration()
		}

		configMarshal, err := json.MarshalIndent(configuration, "", "    ")
		if err != nil {
			printer.GetPrinter().Error("Failed to marshal the configuration.")
			return fmt.Errorf("cmd: failed to marshal the configuration: %w", err)
		}

		printer.GetPrinter().Printf("%s\n", configMarshal)
		return nil
	},
}
