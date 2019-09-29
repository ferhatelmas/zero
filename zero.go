package zero

import "strings"

// Zero width constants
const (
	space             = "\u200B"
	spaceNoBreakSpace = "\uFEFF"
	joiner            = "\u200D"
	nonJoiner         = "\u200C"
)

//nolint: gochecknoglobals
var zeros = []string{
	space,
	spaceNoBreakSpace,
	joiner,
	nonJoiner,
}

// Replace replaces zero width characters with given one
func Replace(s, rep string) string {
	for _, z := range zeros {
		s = strings.ReplaceAll(s, z, rep)
	}
	return s
}

// Has checks if given string has a zero width character
func Has(s string) bool {
	for _, z := range zeros {
		if strings.Contains(s, z) {
			return true
		}
	}
	return false
}
