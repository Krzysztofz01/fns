package utils

import "strings"

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.Title(s)
}

func TrimSelectedWhitespace(b []byte) []byte {
	if len(b) <= 0 {
		return b
	}

	var left int = 0
	for i := 0; i < len(b); i += 1 {
		switch b[i] {
		case '\n':
		case '\r':
			left += 1
		default:
			break
		}
	}

	var right int = len(b)
	for i := len(b) - 1; i > 0; i -= 1 {
		switch b[i] {
		case '\n':
		case '\r':
			right -= 1
		default:
			break
		}
	}

	return b[left:right]
}
