package quicksort

import (
	"math/rand"
	"testing"
	"time"
)

var testData []int

func init() {
	rand.NewSource(time.Now().UnixNano())
	n := 10000000
	// create test data
	testData = make([]int, 0, n)
	for i := 0; i < n; i++ {
		val := rand.Intn(n * 100)
		testData = append(testData, val)
	}
}

func TestQuicksort(t *testing.T) {
	start := time.Now()
	quicksort(testData, 0, len(testData)-1)
	t.Logf("quicksort took %s", time.Since(start))
}

func TestQuicksortConcurrent(t *testing.T) {
	start := time.Now()
	done := make(chan struct{}, 1)
	quicksortConcurrent(testData, 0, len(testData)-1, done)
	<-done
	t.Logf("quicksort concurrent took %s", time.Since(start))
}

func TestQuicksortConcurrentParallel(t *testing.T) {
	start := time.Now()
	done := make(chan struct{}, 1)
	quicksortConcurrentParallel(testData, 0, len(testData)-1, done, 3)
	<-done
	t.Logf("quicksort concurrent took %s", time.Since(start))
}

func BenchmarkQuicksort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quicksort(testData, 0, len(testData)-1)
	}
}

func BenchmarkQuicksortConcurrent(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		done := make(chan struct{}, 1)
		quicksortConcurrent(testData, 0, len(testData)-1, done)
		<-done
	}
}

func BenchmarkQuicksortConcurrentParallel(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		done := make(chan struct{}, 1)
		quicksortConcurrentParallel(testData, 0, len(testData)-1, done, 3)
		<-done
	}
}
