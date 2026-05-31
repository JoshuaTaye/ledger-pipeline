package merchant

import "strings"

// Normalize collapses common merchant description variants.
func Normalize(description string) string {
	d := strings.TrimSpace(description)
	replacements := map[string]string{
		"amzn": "amazon", "amazon.com": "amazon", "uber *trip": "uber",
	}
	for from, to := range replacements {
		if strings.Contains(d, from) {
			return to
		}
	}
	return d
}
