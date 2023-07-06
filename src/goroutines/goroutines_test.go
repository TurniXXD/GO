package goroutines

import "testing"

// circa 0.0000168 ns/op
func BenchmarkSequentialStringBuilder(b *testing.B) {
	sequentialStringBuilder()
}

// circa 0.0000083 ns/op
func BenchmarkConcurrentStringBuilder(b *testing.B) {
	concurrentStringBuilder()
}

// 3002277227 ns/op
// 3002709450 ns/op
// 3002603425 ns/op
// 3002540925 ns/op
// 3003126942 ns/op
func BenchmarkRWMutex(b *testing.B) {
	rw_mutex()
}

// 3002087397 ns/op
// 3002601725 ns/op
// 3001787051 ns/op
// 3001924042 ns/op
// 3001896102 ns/op
func BenchmarkRWMutexSingleChan(b *testing.B) {
	rw_mutex_single_chan()
}
