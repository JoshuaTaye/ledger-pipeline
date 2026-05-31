package ledger

import (
	"fmt"
	"sort"
	"time"

	"github.com/joshuataye/ledgerpipeline/internal/parser"
)

// Entry is a double-entry journal line.
type Entry struct {
	Date        time.Time
	Description string
	Account     string
	Debit       float64
	Credit      float64
}

// Journal holds ordered ledger entries.
type Journal struct {
	entries []Entry
}

func NewJournal() *Journal {
	return &Journal{}
}

func (j *Journal) Post(date time.Time, description, account string, debit, credit float64) {
	j.entries = append(j.entries, Entry{
		Date: date, Description: description, Account: account,
		Debit: debit, Credit: credit,
	})
}

func (j *Journal) FromTransactions(txns []parser.Transaction, cashAccount string) {
	for _, tx := range txns {
		if tx.Amount < 0 {
			j.Post(tx.Date, tx.Description, tx.Category, -tx.Amount, 0)
			j.Post(tx.Date, tx.Description, cashAccount, 0, -tx.Amount)
		} else {
			j.Post(tx.Date, tx.Description, cashAccount, tx.Amount, 0)
			j.Post(tx.Date, tx.Description, tx.Category, 0, tx.Amount)
		}
	}
}

func (j *Journal) Entries() []Entry {
	out := make([]Entry, len(j.entries))
	copy(out, j.entries)
	sort.Slice(out, func(i, k int) bool {
		if out[i].Date.Equal(out[k].Date) {
			return out[i].Account < out[k].Account
		}
		return out[i].Date.Before(out[k].Date)
	})
	return out
}

func (j *Journal) String() string {
	var s string
	for _, e := range j.Entries() {
		s += fmt.Sprintf("%s %s %s dr=%.2f cr=%.2f\n",
			e.Date.Format("2006-01-02"), e.Account, e.Description, e.Debit, e.Credit)
	}
	return s
}
