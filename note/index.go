package note

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/Krzysztofz01/fns/utils"
)

const defaultNoteSliceCapacity = 32

func IndexNotes(ctx context.Context, paths ...string) ([]Note, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("note: no note paths to index specified")
	}

	for _, p := range paths {
		if exist, err := utils.IsDir(p); err != nil {
			return nil, fmt.Errorf("note: directory %s can not be accessed: %w", p, err)
		} else {
			if !exist {
				return nil, fmt.Errorf("note: directory %s does not exist", p)
			}
		}
	}

	noteIndexResults := make([][]Note, len(paths))
	eg, egctx := errgroup.WithContext(ctx)

	for i, p := range paths {
		var (
			index int    = i
			path  string = p
		)

		eg.Go(func() error {
			notes, err := indexNoteDir(egctx, path)
			if err != nil {
				return err
			}

			noteIndexResults[index] = notes
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("note: note indexing was interrupted: %w", err)
		} else {
			return nil, fmt.Errorf("note: note indexing failed: %w", err)
		}
	}

	notes := make([]Note, 0, len(paths)*defaultNoteSliceCapacity)
	for _, resultNotes := range noteIndexResults {
		for _, resultNote := range resultNotes {
			notes = append(notes, resultNote)
		}
	}

	return notes, nil
}

func indexNoteDir(ctx context.Context, p string) ([]Note, error) {
	notes := make([]Note, 0, defaultNoteSliceCapacity)

	err := filepath.Walk(p, func(walkPath string, info fs.FileInfo, walkErr error) error {
		// TODO: Do not fail all if one file access fails
		if walkErr != nil {
			return fmt.Errorf("note: indexing file access failed: %w", walkErr)
		}

		select {
		case <-ctx.Done():
			return io.EOF
		default:
		}

		if info.IsDir() {
			return nil
		}

		if !utils.HasExt(walkPath, targetExtensions) {
			return nil
		}

		// TODO: Do not fail all if one note parsing fails
		note, err := NewNote(walkPath)
		if err != nil {
			return fmt.Errorf("note: failed to create the note: %w", err)
		}

		notes = append(notes, note)
		return nil
	})

	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("note: note indexing was interrupted: %w", err)
		} else {
			return nil, fmt.Errorf("note: note indexing failed: %w", err)
		}
	} else {
		return notes, nil
	}
}
