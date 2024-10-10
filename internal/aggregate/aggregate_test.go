package aggregate

import (
	"testing"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

func tx(category string, amount float64) parser.Transaction {
	return parser.Transaction{
		Date:        time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		Description: "sample",
		Category:    category,
		Amount:      amount,
	}
}

func TestByCategory_GroupsAndTotals(t *testing.T) {
	tests := []struct {
		name string
		txs  []parser.Transaction
		want []CategorySummary
	}{
		{
			name: "single category",
			txs: []parser.Transaction{
				tx("Food", -10),
				tx("Food", -5.50),
			},
			want: []CategorySummary{
				{Category: "Food", TransactionCount: 2, TotalAmount: -15.50, DebitTotal: -15.50},
			},
		},
		{
			name: "mixed debits and credits",
			txs: []parser.Transaction{
				tx("Shopping", -100),
				tx("Shopping", 20),
			},
			want: []CategorySummary{
				{Category: "Shopping", TransactionCount: 2, TotalAmount: -80, DebitTotal: -100, CreditTotal: 20},
			},
		},
		{
			name: "multiple categories sorted",
			txs: []parser.Transaction{
				tx("Transport", -25),
				tx("Food", -12),
				tx("Income", 1000),
			},
			want: []CategorySummary{
				{Category: "Food", TransactionCount: 1, TotalAmount: -12, DebitTotal: -12},
				{Category: "Income", TransactionCount: 1, TotalAmount: 1000, CreditTotal: 1000},
				{Category: "Transport", TransactionCount: 1, TotalAmount: -25, DebitTotal: -25},
			},
		},
		{
			name: "empty category becomes uncategorized",
			txs: []parser.Transaction{
				tx("", -3),
			},
			want: []CategorySummary{
				{Category: "Uncategorized", TransactionCount: 1, TotalAmount: -3, DebitTotal: -3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByCategory(tt.txs)
			if len(got) != len(tt.want) {
				t.Fatalf("len = %d, want %d", len(got), len(tt.want))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Fatalf("summary[%d] = %+v, want %+v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestByCategory_Empty(t *testing.T) {
	got := ByCategory(nil)
	if len(got) != 0 {
		t.Fatalf("len = %d, want 0", len(got))
	}
}

func TestByCategory_SortsAlphabetically(t *testing.T) {
	txs := []parser.Transaction{
		tx("Zebra", -1),
		tx("Alpha", -2),
		tx("Middle", -3),
	}
	got := ByCategory(txs)
	if got[0].Category != "Alpha" || got[1].Category != "Middle" || got[2].Category != "Zebra" {
		t.Fatalf("unexpected order: %+v", got)
	}
}

func TestNetTotal(t *testing.T) {
	tests := []struct {
		name string
		txs  []parser.Transaction
		want float64
	}{
		{name: "empty", txs: nil, want: 0},
		{name: "mixed", txs: []parser.Transaction{tx("A", -10), tx("B", 25)}, want: 15},
		{name: "all debits", txs: []parser.Transaction{tx("A", -5), tx("A", -5)}, want: -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NetTotal(tt.txs); got != tt.want {
				t.Fatalf("NetTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
