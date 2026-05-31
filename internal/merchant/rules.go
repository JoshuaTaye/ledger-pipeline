package merchant

import "strings"

// SuffixRule maps a description suffix to a normalized merchant label.
type SuffixRule struct {
	Suffix string
	Label  string
}

// ApplySuffixRules rewrites descriptions ending with known suffixes.
func ApplySuffixRules(description string, rules []SuffixRule) string {
	normalized := Normalize(description)
	for _, r := range rules {
		suffix := strings.ToLower(strings.TrimSpace(r.Suffix))
		if suffix != "" && strings.HasSuffix(normalized, suffix) {
			return strings.ToLower(strings.TrimSpace(r.Label))
		}
	}
	return normalized
}
