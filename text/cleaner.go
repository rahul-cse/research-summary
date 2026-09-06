package text

import (
	"regexp"
	"strings"
)

// Clean removes common formatting problems from extracted PDF text.
func Clean(input string) string {

	// Normalize line endings.
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")

	// Fix words broken across lines.
	regexHyphenated := regexp.MustCompile(`(\w)-\s*\n\s*(\w)`)
	input = regexHyphenated.ReplaceAllString(input, `$1$2`)

	// Replace tabs with spaces.
	input = strings.ReplaceAll(input, "\t", " ")

	// Remove spaces at the beginning and end of each line.
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	input = strings.Join(lines, "\n")

	// Collapse multiple spaces into one.
	regexSpaces := regexp.MustCompile(`[ ]{2,}`)
	input = regexSpaces.ReplaceAllString(input, " ")

	// Collapse 3 or more consecutive blank lines into 2.
	regexBlankLines := regexp.MustCompile(`\n{3,}`)
	input = regexBlankLines.ReplaceAllString(input, "\n\n")

	// Remove whitespace from the beginning and end of the entire text.
	return strings.TrimSpace(input)
}
