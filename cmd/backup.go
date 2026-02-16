package cmd

import (
	"archive/zip"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path"
	"time"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/note"
	"github.com/Krzysztofz01/fns/printer"
	"github.com/Krzysztofz01/fns/utils"
)

func init() {
	rootCmd.AddCommand(backupCmd)
}

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup all notes",
	Long:  "Create a backup ZIP file with all notes accessible via FNS",
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

		cwd, err := os.Getwd()
		if err != nil {
			printer.GetPrinter().Error("Failed to access the working directory path.")
			return fmt.Errorf("cmd: failed to access the working directory path: %w", err)
		}

		pathSafeTimestamp := time.Now().UTC().Format("2006-01-02T15-04-05")
		backupFileName := fmt.Sprintf("fns-backup-%s.zip", pathSafeTimestamp)
		backupFilePath := path.Join(cwd, backupFileName)

		backupFile, err := os.Create(backupFilePath)
		if err != nil {
			printer.GetPrinter().Error("Failed to create the backup file.")
			return fmt.Errorf("note: failed to create the backup zip file: %w", err)
		}

		defer func() {
			_ = backupFile.Close()
		}()

		zipWriter := zip.NewWriter(backupFile)

		defer func() {
			_ = zipWriter.Flush()
			_ = zipWriter.Close()
		}()

		progressStep, progressFinalize, err := printer.GetPrinter().Progress("Creating notes backup...", len(notes))
		if err != nil {
			return fmt.Errorf("note: failed to print the progress bar: %w", err)
		}

		defer progressFinalize()

		for _, n := range notes {
			if err := writeNoteToZip(n, zipWriter); err != nil {
				printer.GetPrinter().Error("Failed to write the note to the backup file.")
				return fmt.Errorf("note: failed to write not to the backup file: %w", err)
			} else {
				progressStep()
			}
		}

		printer.GetPrinter().Printf("Backup saved at: %s\n", backupFilePath)
		return nil
	},
}

func writeNoteToZip(n note.Note, writer *zip.Writer) error {
	f, err := os.Open(n.GetPath())
	if err != nil {
		return fmt.Errorf("note: failed to open the note file: %w", err)
	}

	defer func() {
		_ = f.Close()
	}()

	zf, err := writer.Create(utils.BaseWithParent(n.GetPath()))
	if err != nil {
		return fmt.Errorf("note: failed to create the note backup entry: %w", err)
	}

	if _, err := io.Copy(zf, f); err != nil {
		return fmt.Errorf("note: failed to copy the note content to the backup entry: %w", err)
	}

	return nil
}
