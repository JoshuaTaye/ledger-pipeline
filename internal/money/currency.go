package money

type Currency string

const (
    USD Currency = "USD"
    EUR Currency = "EUR"
    GBP Currency = "GBP"
)

func (c Currency) Valid() bool {
    return c == USD || c == EUR || c == GBP
}
