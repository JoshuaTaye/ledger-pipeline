package qif

import (
	"strings"
	"testing"
)

func TestParseReader(t *testing.T) {
	raw := "D1/15/2026\nT-9.99\nPCoffee\n^\n"
	r := strings.NewReader(raw)
	txns, err := ParseReader(r)
	if err != nil || len(txns) != 1 {
		t.Fatalf("err=%v len=%d", err, len(txns))
	}
}
