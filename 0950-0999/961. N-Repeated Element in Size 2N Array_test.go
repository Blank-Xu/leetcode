package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array/

func repeatedNTimes(A []int) int {
	m := make(map[int]bool)
	for i := range A {
		if m[A[i]] {
			return A[i]
		}
		m[A[i]] = true
	}
	return 0
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsRepeatedNTimes = []struct {
	input  []int
	output int
}{
	{[]int{1, 2, 3, 3}, 3},
	{[]int{2, 1, 2, 5, 3, 2}, 2},
	{[]int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
}

func Test_repeatedNTimes(t *testing.T) {
	for _, record := range recordsRepeatedNTimes {
		if result := repeatedNTimes(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_repeatedNTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = repeatedNTimes(recordsRepeatedNTimes[0].input)
	}
}
