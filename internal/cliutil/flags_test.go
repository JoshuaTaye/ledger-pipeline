package cliutil

import (
	"flag"
	"testing"
)

func TestParseFilterFlags(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	flags := ParseFilterFlags(fs)
	if err := fs.Parse([]string{"-from", "2026-01-01", "-to", "2026-01-31", "-categories", "Food,Transport"}); err != nil {
		t.Fatal(err)
	}
	opt := flags.Options()
	if opt.From.IsZero() || opt.To.IsZero() {
		t.Fatal("expected date range")
	}
	if len(opt.Categories) != 2 {
		t.Fatalf("Categories = %v", opt.Categories)
	}
}
