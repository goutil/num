// Package num provides some useful algorithms for working with numbers, math,
// etc.
package num

// MinInt64 Returns the integer with the lowest value.
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MaxInt64 Returns the integer with the highest value.
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// AbsInt Returns the absolute value of `v`.
func AbsInt(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

// MinInt Returns the integer with the lowest value.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt Returns the integer with the highest value.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
