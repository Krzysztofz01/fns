package utils

import (
	"path"
	"path/filepath"
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

func BaseWithParent(filePath string) string {
	parent := filepath.Base(filepath.Dir(filePath))
	base := filepath.Base(filePath)

	return path.Join(parent, base)
}

func SplitNameExt(filePath string) (string, string) {
	var (
		base      string   = filepath.Base(filePath)
		baseParts []string = strings.Split(base, ".")
		name      string   = base
		ext       string   = ""
	)

	if len(baseParts) > 0 {
		name = baseParts[0]
	}

	ext = strings.TrimPrefix(base, name)

	return name, ext
}
