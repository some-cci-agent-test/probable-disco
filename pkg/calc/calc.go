package calc

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of a and b.
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of a and b.
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the integer division of a by b. 
// It panics if b is zero.
func Divide(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
