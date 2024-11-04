package account

import "fmt"

// Type classifies a ledger account.
type Type string

const (
	TypeChecking Type = "checking"
	TypeSavings  Type = "savings"
	TypeCredit   Type = "credit"
)

// Account is a bank or card account tracked by the pipeline.
type Account struct {
	ID      string
	Name    string
	Type    Type
	Opening float64
}

func (a Account) Validate() error {
	if a.ID == "" {
		return fmt.Errorf("account id required")
	}
	if a.Type != TypeChecking && a.Type != TypeSavings && a.Type != TypeCredit {
		return fmt.Errorf("unknown account type %q", a.Type)
	}
	return nil
}

// Registry holds accounts keyed by id.
type Registry struct {
	byID map[string]Account
}

func NewRegistry(accounts ...Account) (*Registry, error) {
	r := &Registry{byID: make(map[string]Account, len(accounts))}
	for _, a := range accounts {
		if err := a.Validate(); err != nil {
			return nil, err
		}
		if _, dup := r.byID[a.ID]; dup {
			return nil, fmt.Errorf("duplicate account id %q", a.ID)
		}
		r.byID[a.ID] = a
	}
	return r, nil
}

func (r *Registry) Get(id string) (Account, bool) {
	a, ok := r.byID[id]
	return a, ok
}

func (r *Registry) All() []Account {
	out := make([]Account, 0, len(r.byID))
	for _, a := range r.byID {
		out = append(out, a)
	}
	return out
}
