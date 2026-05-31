package accountsfile

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joshuataye/ledgerpipeline/internal/account"
)

type fileAccount struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Opening float64 `json:"opening"`
}

// Load reads accounts from a JSON array file.
func Load(path string) ([]account.Account, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read accounts: %w", err)
	}
	var raw []fileAccount
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parse accounts: %w", err)
	}
	out := make([]account.Account, 0, len(raw))
	for _, r := range raw {
		out = append(out, account.Account{
			ID: r.ID, Name: r.Name, Type: account.Type(r.Type), Opening: r.Opening,
		})
	}
	return out, nil
}

// LoadRegistry builds a registry from a JSON accounts file.
func LoadRegistry(path string) (*account.Registry, error) {
	accounts, err := Load(path)
	if err != nil {
		return nil, err
	}
	return account.NewRegistry(accounts...)
}
