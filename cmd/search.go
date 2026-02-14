package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/note"
	"github.com/Krzysztofz01/fns/printer"
	"github.com/Krzysztofz01/fns/utils"
)

var (
	PrintPath bool
)

func init() {
	searchCmd.PersistentFlags().BoolVarP(&PrintPath, "path", "p", false, "Print the file path instead of its content.")

	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search and print a note",
	Long:  "Perform a fuzzy search of indexed notes and print the content of the selected one",
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
			return fmt.Errorf("cmd: failed to perform the fuzzy note select: %w", err)
		}

		note, ok := reverseLookup[selected]
		if !ok {
			printer.GetPrinter().Error("Failed to access the selected note.")
			return fmt.Errorf("cmd: the selected note is not present in the lookup")
		}

		if PrintPath {
			printer.GetPrinter().Printf("%s\n", note.GetPath())
			return nil
		}

		file, err := os.Open(note.GetPath())
		if err != nil {
			printer.GetPrinter().Error("Failed to open the selected note.")
			return fmt.Errorf("cmd: failed to open the selected note file: %w", err)
		}

		defer func() {
			_ = file.Close()
		}()

		noteContent, err := io.ReadAll(file)
		if err != nil {
			printer.GetPrinter().Error("Failed to read the note file content.")
			return fmt.Errorf("cmd: failed to read the opened note file content: %w", err)
		}

		if config.GetConfiguration().TrimNote {
			noteContent = utils.TrimSelectedWhitespace(noteContent)
		}

		printer.GetPrinter().Printf("%s\n", noteContent)
		return nil
	},
}
