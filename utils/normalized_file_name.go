package utils

import (
	"regexp"
	"strings"
)

// NormalizedFileName as Arma modders do not follow naming conventions, it is necessary to standardize file names.
func NormalizedFileName(fileName string) string {
	fileName = regexp.MustCompile(`[(?:\-)]`).ReplaceAllString(fileName, "_")
	fileName = regexp.MustCompile(`[^a-zA-Z0-9_.]`).ReplaceAllString(fileName, "")
	fileName = strings.ToLower(fileName)

	return fileName
}
