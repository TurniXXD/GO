package snippets

import (
	"testing"
)

// circa 0.0002417 ns/op
func BenchmarkConcateStrings(b *testing.B) {
	ConcateStrings("test")
}

// circa 0.0000108 ns/op
func BenchmarkConcateStringsWithBuffer(b *testing.B) {
	ConcateStringsWithBuffer("test")
}

// circa 0.0000082 ns/op
func BenchmarkConcateStringsWithStringBuilder(b *testing.B) {
	ConcateStringsWithStringBuilder("test")
}
