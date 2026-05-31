package merchant

import "testing"

func TestApplySuffixRules(t *testing.T) {
	rules := []SuffixRule{
		{Suffix: " inc", Label: "Acme"},
		{Suffix: " llc", Label: "Widgets"},
	}
	got := ApplySuffixRules("ACME INC", rules)
	if got != "acme" {
		t.Fatalf("ApplySuffixRules() = %q", got)
	}
	got = ApplySuffixRules("Widgets LLC", rules)
	if got != "widgets" {
		t.Fatalf("ApplySuffixRules() = %q", got)
	}
}
