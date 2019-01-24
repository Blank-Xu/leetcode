package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	n := 0
	for _, v := range nums {
		if v != nums[n] {
			n++
			nums[n] = v
		}
	}
	return n + 1
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsRemoveDuplicates = []struct {
	input  []int
	output int
}{
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func Test_removeDuplicates(t *testing.T) {
	for _, record := range recordsRemoveDuplicates {
		if result := removeDuplicates(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeDuplicates(recordsRemoveDuplicates[1].input)
	}
}
