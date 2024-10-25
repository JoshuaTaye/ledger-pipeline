package stages

import "github.com/joshuataye/ledgerpipeline/internal/parser"

type Registry struct {
    list []Stage
}

func NewRegistry(st ...Stage) *Registry {
    return &Registry{list: st}
}

func (r *Registry) Run(txns []parser.Transaction) ([]parser.Transaction, error) {
    cur := txns
    for _, st := range r.list {
        next, err := st.Run(cur)
        if err != nil {
            return nil, err
        }
        cur = next
    }
    return cur, nil
}
