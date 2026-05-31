package cliutil

import (
	"flag"
	"strings"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/filter"
)

// FilterFlags holds registered filter flag pointers.
type FilterFlags struct {
	from       *string
	to         *string
	categories *string
	minAmount  *float64
	maxAmount  *float64
}

// ParseFilterFlags registers filter flags on fs and returns accessors.
func ParseFilterFlags(fs *flag.FlagSet) FilterFlags {
	return FilterFlags{
		from:       fs.String("from", "", "include transactions on or after date (YYYY-MM-DD)"),
		to:         fs.String("to", "", "include transactions on or before date (YYYY-MM-DD)"),
		categories: fs.String("categories", "", "comma-separated category allowlist"),
		minAmount:  fs.Float64("min-amount", 0, "minimum transaction amount"),
		maxAmount:  fs.Float64("max-amount", 0, "maximum transaction amount"),
	}
}

// Options builds filter.Options from parsed flag values.
func (f FilterFlags) Options() filter.Options {
	opt := filter.Options{MinAmount: *f.minAmount, MaxAmount: *f.maxAmount}
	if *f.from != "" {
		if t, err := time.Parse("2006-01-02", *f.from); err == nil {
			opt.From = t
		}
	}
	if *f.to != "" {
		if t, err := time.Parse("2006-01-02", *f.to); err == nil {
			opt.To = t
		}
	}
	if *f.categories != "" {
		opt.Categories = map[string]struct{}{}
		for _, c := range strings.Split(*f.categories, ",") {
			c = strings.TrimSpace(c)
			if c != "" {
				opt.Categories[c] = struct{}{}
			}
		}
	}
	return opt
}
