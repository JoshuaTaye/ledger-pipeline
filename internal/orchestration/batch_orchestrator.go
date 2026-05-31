package orchestration

import (
	"github.com/joshuataye/ledgerpipeline/internal/account"
	"github.com/joshuataye/ledgerpipeline/internal/budget"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
	"github.com/joshuataye/ledgerpipeline/internal/pipeline"
)

// RunConfig drives a full ledger pipeline quote-style run.
type RunConfig struct {
	AccountID    string
	Pipeline     pipeline.Config
	BudgetLines  []budget.Line
	Accounts     *account.Registry
}

// RunResult aggregates pipeline and budget outputs.
type RunResult struct {
	Pipeline pipeline.Result
	Budget   []budget.Variance
	Account  account.Account
}

// BatchOrchestrator coordinates account lookup, pipeline execution, and budget comparison.
type BatchOrchestrator struct {
	accounts *account.Registry
}

func NewBatchOrchestrator(accounts *account.Registry) *BatchOrchestrator {
	return &BatchOrchestrator{accounts: accounts}
}

func (o *BatchOrchestrator) Run(txns []parser.Transaction, cfg RunConfig) (RunResult, error) {
	acct, ok := o.accounts.Get(cfg.AccountID)
	if !ok {
		return RunResult{}, parser.ErrInvalidInputMsg("unknown account " + cfg.AccountID)
	}
	pcfg := cfg.Pipeline
	if pcfg.Reconcile == nil && acct.Opening != 0 {
		pcfg.Reconcile = &pipeline.ReconcileConfig{Opening: 0}
	}
	result, err := pipeline.Run(txns, pcfg)
	if err != nil {
		return RunResult{}, err
	}
	actual := map[string]float64{}
	for _, s := range result.Summaries {
		actual[s.Category] = s.TotalAmount
	}
	var variances []budget.Variance
	if len(cfg.BudgetLines) > 0 {
		variances = budget.Compare(actual, cfg.BudgetLines)
	}
	return RunResult{
		Pipeline: result,
		Budget:   variances,
		Account:  acct,
	}, nil
}
