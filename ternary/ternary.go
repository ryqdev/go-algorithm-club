package ternary

// Ternary is the conditional operator in Golang, the updated version of IfThenElse
// Example:
// x, y := 1, 2
// max, ok := Ternary(x > y, x, y).(int)
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
