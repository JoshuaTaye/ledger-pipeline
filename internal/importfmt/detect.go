package importfmt

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/joshuataye/ledgerpipeline/internal/import/fixedwidth"
	"github.com/joshuataye/ledgerpipeline/internal/import/ofx"
	"github.com/joshuataye/ledgerpipeline/internal/import/qif"
	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Format identifies a supported import layout.
type Format string

const (
	FormatCSV        Format = "csv"
	FormatOFX        Format = "ofx"
	FormatQIF        Format = "qif"
	FormatFixedWidth Format = "fixedwidth"
)

// DefaultFixedWidthLayout returns the standard bank export column layout.
func DefaultFixedWidthLayout() fixedwidth.Layout {
	return fixedwidth.Layout{
		Date:        fixedwidth.Column{0, 10},
		Description: fixedwidth.Column{10, 30},
		Amount:      fixedwidth.Column{30, 40},
	}
}

// Detect inspects a file path and returns its format.
func Detect(path string) (Format, error) {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".csv":
		return FormatCSV, nil
	case ".ofx":
		return FormatOFX, nil
	case ".qif":
		return FormatQIF, nil
	case ".txt":
		return FormatFixedWidth, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := string(data)
	if strings.Contains(text, "<OFX>") {
		return FormatOFX, nil
	}
	if strings.Contains(text, "!Type:") {
		return FormatQIF, nil
	}
	if strings.Contains(text, "date,description") {
		return FormatCSV, nil
	}
	return FormatFixedWidth, nil
}

// ParseFile loads transactions using detected format.
func ParseFile(path string) ([]parser.Transaction, error) {
	format, err := Detect(path)
	if err != nil {
		return nil, err
	}
	switch format {
	case FormatCSV:
		return parser.ParseFile(path)
	case FormatOFX:
		return ofx.ParseFile(path)
	case FormatQIF:
		return qif.ParseFile(path)
	case FormatFixedWidth:
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		lines := strings.Split(string(data), "\n")
		return fixedwidth.ParseLines(lines, DefaultFixedWidthLayout())
	default:
		return parser.ParseFile(path)
	}
}
