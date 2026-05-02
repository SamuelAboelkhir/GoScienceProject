// Package utils: A utility package with helper functions
package utils

import "strings"

func CleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}
