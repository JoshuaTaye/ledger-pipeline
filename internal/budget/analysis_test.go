package budget

import "testing"

func TestAnalyze(t *testing.T) {
	variances := []Variance{
		{Category: "Food", Limit: 200, Actual: -150, Delta: 50},
	}
	util := Analyze(variances)
	if len(util) != 1 || util[0].UsedPct != 75 {
		t.Fatalf("Analyze() = %+v", util)
	}
}

func TestOverBudget(t *testing.T) {
	variances := []Variance{
		{Category: "Food", Limit: 100, Actual: -150},
		{Category: "Transport", Limit: 50, Actual: -20},
	}
	over := OverBudget(variances)
	if len(over) != 1 || over[0].Category != "Food" {
		t.Fatalf("OverBudget() = %+v", over)
	}
}
