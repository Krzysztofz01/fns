package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/editor"
	"github.com/Krzysztofz01/fns/note"
	"github.com/Krzysztofz01/fns/printer"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Search and edit a note",
	Long:  "Perform a fuzzy search of indexed notes and edit the content of the selected one with the editor specified in the configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		noteDirPaths := config.GetConfiguration().NoteReadDirectoryPaths
		if len(noteDirPaths) <= 0 {
			printer.GetPrinter().Warn("Configure the note source directory paths via 'note-read-directory-paths' config entry.")
			return nil
		}

		notes, err := note.IndexNotes(cmd.Context(), noteDirPaths...)
		if err != nil {
			printer.GetPrinter().Error("Failed to access notes.")
			return fmt.Errorf("cmd: failed to index the notes: %w", err)
		}

		var (
			options       []string             = make([]string, 0, len(notes))
			reverseLookup map[string]note.Note = make(map[string]note.Note, len(notes))
		)

		for _, note := range notes {
			options = append(options, note.GetName())
			reverseLookup[note.GetName()] = note
		}

		selected, err := printer.GetPrinter().FuzzySelect("Indexed notes", options)
		if err != nil {
			printer.GetPrinter().Error("Failed to search across notes.")
			return fmt.Errorf("cmd: failed to perform the fuzzy note edit select: %w", err)
		}

		note, ok := reverseLookup[selected]
		if !ok {
			printer.GetPrinter().Error("Failed to access the selected note.")
			return fmt.Errorf("cmd: the selected for edit note is not present in the lookup")
		}

		_, err = editor.ExecEditor(cmd.Context(), note.GetPath())
		if err != nil {
			printer.GetPrinter().Error("Failed to edit the note at %s.", note.GetPath())
			return fmt.Errorf("cmd: failed to edit the note: %w", err)
		}

		return nil
	},
}
