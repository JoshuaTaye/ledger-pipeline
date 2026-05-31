package filter

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/config"
)

func TestFromConfig(t *testing.T) {
	opt := FromConfig(config.File{
		MinDate: "2026-01-01",
		MaxDate: "2026-01-31",
		Categories: []string{"Food", "Transport"},
	})
	if opt.From.IsZero() || opt.To.IsZero() || len(opt.Categories) != 2 {
		t.Fatalf("FromConfig() = %+v", opt)
	}
}

func TestPresetLastNDays(t *testing.T) {
	end := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC)
	opt := PresetLastNDays(end, 7)
	if opt.From.Day() != 4 || opt.To != end {
		t.Fatalf("PresetLastNDays() = %+v", opt)
	}
}
