package pipeline

import (
	"fmt"

	"github.com/joshuataye/ledgerpipeline/internal/audit"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// AuditResult wraps a pipeline result with an audit log.
type AuditResult struct {
	Result Result
	Log    *audit.Log
}

// RunWithAudit executes the pipeline and records stage events.
// Transactions without descriptions are rejected.
func RunWithAudit(txns []parser.Transaction, cfg Config) (AuditResult, error) {
	log := audit.NewLog()
	log.Record("start", fmt.Sprintf("processing %d transactions", len(txns)))
	for i, tx := range txns {
		if tx.Description == "" {
			return AuditResult{}, fmt.Errorf("transaction %d: description required", i)
		}
	}
	result, err := Run(txns, cfg)
	if err != nil {
		log.Record("error", err.Error())
		return AuditResult{Log: log}, err
	}
	log.Record("complete", fmt.Sprintf("net total %.2f", result.NetTotal))
	return AuditResult{Result: result, Log: log}, nil
}
