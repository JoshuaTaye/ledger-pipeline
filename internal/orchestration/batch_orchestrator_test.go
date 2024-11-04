package orchestration

import (
	"testing"

	"github.com/joshuataye/ledgerpipeline/internal/account"
	"github.com/joshuataye/ledgerpipeline/internal/budget"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
	"github.com/joshuataye/ledgerpipeline/internal/pipeline"
	"github.com/joshuataye/ledgerpipeline/internal/validate"
)

func TestBatchOrchestratorRun(t *testing.T) {
	reg, err := account.NewRegistry(account.Account{
		ID: "main", Name: "Checking", Type: account.TypeChecking, Opening: 500,
	})
	if err != nil {
		t.Fatal(err)
	}
	orch := NewBatchOrchestrator(reg)
	txns := []parser.Transaction{
		{Description: "Coffee", Category: "Food", Amount: -4.5},
		{Description: "Pay", Category: "Income", Amount: 100},
	}
	result, err := orch.Run(txns, RunConfig{
		AccountID: "main",
		Pipeline: pipeline.Config{
			Validate: validate.Validator{},
		},
		BudgetLines: []budget.Line{{Category: "Food", Limit: 200}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Budget) != 1 || result.Account.ID != "main" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestBatchOrchestratorUnknownAccount(t *testing.T) {
	reg, _ := account.NewRegistry()
	orch := NewBatchOrchestrator(reg)
	_, err := orch.Run(nil, RunConfig{AccountID: "missing"})
	if err == nil {
		t.Fatal("expected error")
	}
}
