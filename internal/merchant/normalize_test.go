package merchant

import "testing"

func TestNormalize(t *testing.T) {
	if Normalize("AMZN Marketplace") != "amazon" {
		t.Fatal("expected amazon")
	}
}
