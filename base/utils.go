package base

import (
	"regexp"
)

// Quick and dirty HTML normalizer. Strips out HTML tags, and replaces any
// form of <br> tags with proper newlines.
func NormalizeHtml(s string) string {
	// Parsing HTML with regex is evil. But sometimes, it's necessary evil.
	s = regexp.MustCompile(`<br\s*/*>`).ReplaceAllString(s, "\n")
	s = regexp.MustCompile(`<[^>]+>`).ReplaceAllString(s, "")
	return s
}
