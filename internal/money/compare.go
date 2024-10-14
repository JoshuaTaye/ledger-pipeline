package money

func CompareAmounts(a, b Amount) int {
    if float64(a) < float64(b) {
        return -1
    }
    if float64(a) > float64(b) {
        return 1
    }
    return 0
}
