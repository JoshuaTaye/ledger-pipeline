package stages

import (
    "github.com/joshuataye/ledgerpipeline/internal/parser"
    "github.com/joshuataye/ledgerpipeline/internal/dedupe"
    "github.com/joshuataye/ledgerpipeline/internal/filter"
    "github.com/joshuataye/ledgerpipeline/internal/validate"
)

type Stage interface {
    Name() string
    Run([]parser.Transaction) ([]parser.Transaction, error)
}

type DedupeStage struct{}

func (DedupeStage) Name() string { return "dedupe" }
func (DedupeStage) Run(txns []parser.Transaction) ([]parser.Transaction, error) {
    return dedupe.RemoveDuplicates(txns), nil
}

type FilterStage struct{ Opts filter.Options }

func (s FilterStage) Name() string { return "filter" }
func (s FilterStage) Run(txns []parser.Transaction) ([]parser.Transaction, error) {
    return filter.Apply(txns, s.Opts), nil
}

type ValidateStage struct{ V validate.Validator }

func (s ValidateStage) Name() string { return "validate" }
func (s ValidateStage) Run(txns []parser.Transaction) ([]parser.Transaction, error) {
    return txns, s.V.ValidateAll(txns)
}
