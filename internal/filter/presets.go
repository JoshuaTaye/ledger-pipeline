package filter

import (
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/config"
)

// FromConfig converts a JSON config file into filter options.
func FromConfig(cfg config.File) Options {
	opt := Options{}
	if cfg.MinDate != "" {
		if t, err := time.Parse("2006-01-02", cfg.MinDate); err == nil {
			opt.From = t
		}
	}
	if cfg.MaxDate != "" {
		if t, err := time.Parse("2006-01-02", cfg.MaxDate); err == nil {
			opt.To = t
		}
	}
	if len(cfg.Categories) > 0 {
		opt.Categories = make(map[string]struct{}, len(cfg.Categories))
		for _, c := range cfg.Categories {
			opt.Categories[c] = struct{}{}
		}
	}
	return opt
}

// PresetLastNDays returns a filter for the trailing n calendar days ending at end.
func PresetLastNDays(end time.Time, n int) Options {
	if n <= 0 {
		return Options{}
	}
	start := end.AddDate(0, 0, -(n - 1))
	return Options{From: start, To: end}
}
