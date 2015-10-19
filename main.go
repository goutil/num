package num

// MinInt64 Returns the integer with the lowest value.
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}
