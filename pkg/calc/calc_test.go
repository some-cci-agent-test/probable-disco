package calc

import (
	"math/rand"
	"os"
	"testing"
)

// TestAdd is a stable test for Add. 
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %d; want 5", got)
	}
}

// TestSubtract is a stable test for Subtract.
func TestSubtract(t *testing.T) {
	if got := Subtract(10, 6); got != 4 {
		t.Fatalf("Subtract(10,6) = %d; want 4", got)
	}
}

// TestMultiply is a stable test for Multiply.
func TestMultiply(t *testing.T) {
	if got := Multiply(7, 8); got != 56 {
		t.Fatalf("Multiply(7,8) = %d; want 56", got)
	}
}

// TestDivideByZero panics as expected.
func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Divide did not panic on division by zero")
		}
	}()
	_ = Divide(1, 0)
}

// TestFlaky simulates flakiness unless FLAKY_OFF is set.
// Behavior:
// - If env FLAKY_OFF is set to any non-empty value, the test is stable and passes.
// - Otherwise, it fails randomly ~30% of the time.
func TestFlaky(t *testing.T) {
	if os.Getenv("FLAKY_OFF") != "" {
		return
	}

	if rand.Float64() < 0.3 { // 30% chance to fail
		t.Fatalf("flaky failure: random threshold hit")
	}
}
