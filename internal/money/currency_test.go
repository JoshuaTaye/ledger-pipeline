package money

import "testing"

func TestCurrencyValid(t *testing.T) {
    if !USD.Valid() || EUR.Valid() == false {
        t.Fatal("expected valid currencies")
    }
}
