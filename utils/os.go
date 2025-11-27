package utils

import (
	"fmt"
	"os"
)

func IsDir(path string) (bool, error) {
	if stat, err := os.Stat(path); err != nil {
		return false, fmt.Errorf("utils: failed to check if path is pointing to dir: %w", err)
	} else {
		return stat.IsDir(), nil
	}
}

func FileExist(path string) (bool, error) {
	if stat, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, fmt.Errorf("utils: failed to determine if the file exists: %w", err)
		}
	} else {
		if !stat.IsDir() {
			return true, nil
		} else {
			return false, fmt.Errorf("utils: the given path is not pointing to a file: %w", err)
		}
	}
}
