package pipeline

import (
	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
	"github.com/joshuataye/ledgerpipeline/internal/categorize/rules"
	"github.com/joshuataye/ledgerpipeline/internal/dedupe"
	"github.com/joshuataye/ledgerpipeline/internal/filter"
	"github.com/joshuataye/ledgerpipeline/internal/merchant"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
	"github.com/joshuataye/ledgerpipeline/internal/reconcile"
	"github.com/joshuataye/ledgerpipeline/internal/recurring"
	"github.com/joshuataye/ledgerpipeline/internal/tags"
	"github.com/joshuataye/ledgerpipeline/internal/validate"
)

// Config controls batch processing stages.
type Config struct {
	Filter       filter.Options
	Validate     validate.Validator
	Dedupe       bool
	Normalize      bool
	CategorizeRules []rules.Rule
	TagRules       map[string]string
	MinRecurring   int
	Reconcile    *ReconcileConfig
}

// ReconcileConfig enables opening/closing balance checks.
type ReconcileConfig struct {
	Opening float64
	Closing float64
}

// Result is the output of a full batch pipeline run.
type Result struct {
	Transactions []parser.Transaction
	Summaries    []aggregate.CategorySummary
	NetTotal     float64
	Reconcile    *reconcile.Result
	Recurring    []recurring.Candidate
}

func Run(txns []parser.Transaction, cfg Config) (Result, error) {
	preFilter := txns
	if cfg.Dedupe {
		txns = dedupe.RemoveDuplicates(txns)
	}
	txns = filter.Apply(txns, cfg.Filter)
	if cfg.Normalize {
		for i := range txns {
			txns[i].Description = merchant.Normalize(txns[i].Description)
		}
	}
	if len(cfg.CategorizeRules) > 0 {
		txns = rules.Apply(txns, cfg.CategorizeRules)
	}
	if len(cfg.TagRules) > 0 {
		txns = tags.Enrich(txns, cfg.TagRules)
	}
	if err := cfg.Validate.ValidateAll(txns); err != nil {
		return Result{}, err
	}
	summaries := aggregate.ByCategory(txns)
	result := Result{
		Transactions: txns,
		Summaries:    summaries,
		NetTotal:     aggregate.NetTotal(preFilter),
	}
	if cfg.Reconcile != nil {
		r := reconcile.Reconcile(cfg.Reconcile.Opening, cfg.Reconcile.Closing, txns)
		result.Reconcile = &r
	}
	if cfg.MinRecurring > 0 {
		result.Recurring = recurring.Detect(txns, cfg.MinRecurring)
	}
	return result, nil
}
