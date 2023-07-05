package functions

import (
	"math/big"
	"testing"
)

// Test function that is access with native go testing tool
// Its name needs to start with the keyword "Test"
// Its input argument is then the testing output
func TestFibInt(t *testing.T) {
	result := 832040
	if FibInt(30) != result {
		t.Error("Incorrect!")
	}
}

func TestFibBig(t *testing.T) {
	result := big.NewInt(832040)
	if result.Cmp(FibBig(30)) != 0 {
		t.Error("Incorrect!")
	}
}

// Benchmark are often used alongside the testing functions
// Its name needs to start with the keyword "Benchmark"
// Its input argument is then the benchmark output

func BenchmarkFibBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibBig(30)
	}
}

func BenchmarkFibInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibInt(30)
	}
}
