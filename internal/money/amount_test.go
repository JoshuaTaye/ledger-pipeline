package money

import "testing"

func TestParseAmount(t *testing.T) {
	a, err := ParseAmount("-12.50")
	if err != nil || a.Float64() != -12.5 {
		t.Fatalf("got %v err %v", a, err)
	}
}
