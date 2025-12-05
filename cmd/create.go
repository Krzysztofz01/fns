package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/editor"
	"github.com/Krzysztofz01/fns/printer"
	"github.com/Krzysztofz01/fns/utils"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [filename]",
	Short: "Create a new note with the given filename",
	Long:  "Create a new note at the path specified in the configuration and open it with the editor specified in the configuration",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(config.GetConfiguration().NoteWriteDirectoryPath) <= 0 {
			printer.GetPrinter().Warn("Configure the note destination directory via 'note-write-directory-path' config entry.")
			return nil
		}

		var filename string
		if len(args) == 0 {
			var err error
			if filename, err = printer.GetPrinter().TextInput("File name (with extension):"); err != nil {
				printer.GetPrinter().Error("Failed to provide the note file name")
				return fmt.Errorf("cmd: failed to access the file name via text input: %w", err)
			}
		} else {
			filename = args[0]
		}

		if len(filename) <= 0 {
			printer.GetPrinter().Error("The provided note file path is invalid.")
			return fmt.Errorf("cmd: invalid empty note file name")
		}

		notePath := path.Join(config.GetConfiguration().NoteWriteDirectoryPath, filename)

		if exist, err := utils.FileExist(notePath); err != nil {
			printer.GetPrinter().Error("Can not check if the given note file already exist.")
			return fmt.Errorf("cmd: target file exist check failed: %w", err)
		} else if exist {
			printer.GetPrinter().Warn("Note with the given name already exist. Note will not be created.")
			return nil
		}

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
