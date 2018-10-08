package safekubectl

import (
	"os"
	"github.com/fatih/color"
)

// Returns whether the given file or directory exists.
func fileExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Returns the common color used for text highlighting
func GetHighlightColor() (*color.Color) {
	return color.New(color.FgCyan).Add(color.Bold)
}
