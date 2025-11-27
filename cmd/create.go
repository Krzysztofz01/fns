package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/editor"
	"github.com/Krzysztofz01/fns/printer"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [filename]",
	Short: "Create a new note with the given filename",
	Long:  "Create a new note at the path specified in the configuration and open it with the editor specified in the configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(config.GetConfiguration().NoteWriteDirectoryPath) <= 0 {
			printer.GetPrinter().Warn("Configure the note destination directory via 'note-write-directory-path' config entry.")
			return nil
		}

		filename := args[0]
		if len(filename) <= 0 {
			printer.GetPrinter().Error("You must specify the note filename.")
			return fmt.Errorf("cmd: invalid empty note file name")
		}

		notePath := path.Join(config.GetConfiguration().NoteWriteDirectoryPath, filename)

		created, err := editor.ExecEditor(cmd.Context(), notePath)
		if err != nil {
			printer.GetPrinter().Error("Failed to create the note %s at %s.", filename, notePath)
			return fmt.Errorf("cmd: failed to create the note: %w", err)
		}

		if created {
			printer.GetPrinter().Info("Note created at: %s successfully.", notePath)
		}

		return nil
	},
}
