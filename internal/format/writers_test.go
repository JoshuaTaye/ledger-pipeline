package format

import (
	"bytes"
	"strings"
	"testing"

	"github.com/joshuataye/ledgerpipeline/internal/aggregate"
)

func TestWriteCSV(t *testing.T) {
	var buf bytes.Buffer
	summaries := []aggregate.CategorySummary{
		{Category: "Food", TransactionCount: 2, TotalAmount: -50, DebitTotal: -50},
	}
	if err := WriteCSV(&buf, summaries); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "Food,2,-50.00") {
		t.Fatalf("WriteCSV() = %q", buf.String())
	}
}

func TestWriteTSV(t *testing.T) {
	var buf bytes.Buffer
	summaries := []aggregate.CategorySummary{
		{Category: "Food", TransactionCount: 1, TotalAmount: -10},
	}
	if err := WriteTSV(&buf, summaries); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "Food\t1\t-10.00") {
		t.Fatalf("WriteTSV() = %q", buf.String())
	}
}

func TestWriteMarkdown(t *testing.T) {
	var buf bytes.Buffer
	summaries := []aggregate.CategorySummary{
		{Category: "Food", TransactionCount: 1, TotalAmount: -10},
	}
	if err := WriteMarkdown(&buf, summaries, -10); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "| Food |") {
		t.Fatalf("WriteMarkdown() = %q", buf.String())
	}
}
