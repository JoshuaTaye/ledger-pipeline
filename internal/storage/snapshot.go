package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

type snapshotTxn struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
}

// Snapshot is a persisted transaction batch.
type Snapshot struct {
	Created      time.Time `json:"created"`
	Transactions []snapshotTxn `json:"transactions"`
}

// Save writes transactions to a JSON snapshot file.
func Save(path string, txns []parser.Transaction) error {
	snap := Snapshot{
		Created: time.Now().UTC(),
	}
	for _, tx := range txns {
		snap.Transactions = append(snap.Transactions, snapshotTxn{
			Date: tx.Date.Format("2006-01-02"),
			Description: tx.Description,
			Category: tx.Category,
			Amount: tx.Amount,
		})
	}
	data, err := json.MarshalIndent(snap, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// Load reads transactions from a JSON snapshot file.
func Load(path string) ([]parser.Transaction, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read snapshot: %w", err)
	}
	var snap Snapshot
	if err := json.Unmarshal(data, &snap); err != nil {
		return nil, fmt.Errorf("parse snapshot: %w", err)
	}
	out := make([]parser.Transaction, 0, len(snap.Transactions))
	for _, st := range snap.Transactions {
		d, err := time.Parse("2006-01-02", st.Date)
		if err != nil {
			return nil, err
		}
		out = append(out, parser.Transaction{
			Date: d, Description: st.Description, Category: st.Category, Amount: st.Amount,
		})
	}
	return out, nil
}
