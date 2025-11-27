package utils

import (
	"path"
	"strings"
)

func HasExt(filePath string, extensions []string) bool {
	fileExt := strings.ToLower(path.Ext(filePath))

	for _, ext := range extensions {
		if fileExt == strings.ToLower(ext) {
			return true
		}
	}

	return false
}
