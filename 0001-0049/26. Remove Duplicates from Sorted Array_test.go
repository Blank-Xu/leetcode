package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

// Runtime: 8 ms, faster than 89.85% of Go online submissions for Remove Duplicates from Sorted Array.
// Memory Usage: 4.6 MB, less than 75.38% of Go online submissions for Remove
func removeDuplicates(nums []int) int {
	n := 0

	for idx := range nums {
		if nums[idx] != nums[n] {
			n++
			nums[n] = nums[idx]
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
	for idx := range recordsRemoveDuplicates {
		for i := 0; i < b.N; i++ {
			_ = removeDuplicates(recordsRemoveDuplicates[idx].input)
		}
	}
}
