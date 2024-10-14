package money

import "testing"

func TestCompareAmounts(t *testing.T) {
    if CompareAmounts(Amount(1), Amount(2)) != -1 {
        t.Fatal("expected -1")
    }
}
