package budget

import "testing"

func TestCompare(t *testing.T) {
	v := Compare(map[string]float64{"Food": -120}, []Line{{Category: "Food", Limit: -200}})
	if v[0].Delta != -320 {
		t.Fatalf("delta %v", v[0].Delta)
	}
}
