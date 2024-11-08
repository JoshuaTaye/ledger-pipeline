package fixedwidth

import "testing"

func TestParseLine(t *testing.T) {
	layout := Layout{
		Date:        Column{0, 10},
		Description: Column{10, 30},
		Amount:      Column{30, 40},
	}
	line := "2026-01-05Grocery             -12.50    "
	tx, err := ParseLine(line, layout)
	if err != nil || tx.Amount != -12.5 {
		t.Fatalf("err=%v tx=%+v", err, tx)
	}
}
