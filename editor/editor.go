package editor

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/Krzysztofz01/fns/config"
	"github.com/Krzysztofz01/fns/utils"
)

const defaultEditorEnvVariable string = "EDITOR"

func ExecEditor(ctx context.Context, path string) (bool, error) {
	configuration := config.GetConfiguration()

	editorPath := os.Getenv(defaultEditorEnvVariable)
	if len(configuration.EditorPath) > 0 {
		editorPath = configuration.EditorPath
	}

	if len(editorPath) <= 0 {
		return false, fmt.Errorf("editor: no editor available")
	}

	cmd := exec.CommandContext(ctx, editorPath, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	created, _ := utils.FileExist(path)

	if err != nil {
		return created, fmt.Errorf("editor: external editor process call failed: %w", err)
	} else {
		return created, nil
	}
}
