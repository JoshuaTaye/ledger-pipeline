package money

import (
	"fmt"
	"math"
)

// Amount is a statement line amount in major currency units.
type Amount float64

func (a Amount) Float64() float64 { return float64(a) }

func (a Amount) Add(b Amount) Amount { return Amount(float64(a) + float64(b)) }

func (a Amount) IsDebit() bool { return float64(a) < 0 }

func (a Amount) Abs() Amount {
	v := float64(a)
	if v < 0 {
		return Amount(-v)
	}
	return a
}

func ParseAmount(s string) (Amount, error) {
	var v float64
	_, err := fmt.Sscanf(s, "%f", &v)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return 0, fmt.Errorf("invalid amount")
	}
	return Amount(v), nil
}
