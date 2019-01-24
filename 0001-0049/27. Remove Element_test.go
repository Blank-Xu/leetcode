package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	n := -1
	for i := 0; i < l; i++ {
		if nums[i] != val {
			n++
			nums[n] = nums[i]
		}
	}
	return n + 1
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsRemoveElement = []struct {
	nums   []int
	val    int
	output int
}{
	{[]int{3, 2, 2, 3}, 3, 2},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
}

func Test_removeElement(t *testing.T) {
	for _, record := range recordsRemoveElement {
		if result := removeElement(record.nums, record.val); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_removeElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = removeElement(recordsRemoveElement[0].nums, recordsRemoveElement[0].val)
	}
}
