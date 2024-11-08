package ofx

import "testing"

func TestParseSnippet(t *testing.T) {
    text := "DTPOSTED:20260101\nTRNAMT:-4.50\nNAME:Coffee\n"
    txns, err := ParseSnippet(text)
    if err != nil || len(txns) != 1 || txns[0].Amount != -4.5 {
        t.Fatalf("err=%v txns=%+v", err, txns)
    }
}
