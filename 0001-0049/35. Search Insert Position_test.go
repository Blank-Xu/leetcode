package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/search-insert-position/

func searchInsert1(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	index := -1
	for i := 0; i < l; i++ {
		if nums[i] < target {
			continue
		}
		index = i
		break
	}

	if target > nums[l-1] {
		index = l
	}

	return index
}

// Runtime: 4 ms, faster than 89.86% of Go online submissions for Search Insert Position.
// Memory Usage: 3.1 MB, less than 99.76% of Go online submissions for Search Insert Position.
func searchInsert2(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	for i := 0; i < l; i++ {
		if nums[i] < target {
			continue
		}

		return i
	}

	return l
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsSearchInsert = []struct {
	nums   []int
	val    int
	output int
}{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
}

func Test_searchInsert1(t *testing.T) {
	for _, record := range recordsSearchInsert {
		if result := searchInsert1(record.nums, record.val); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Test_searchInsert2(t *testing.T) {
	for _, record := range recordsSearchInsert {
		if result := searchInsert2(record.nums, record.val); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_searchInsert1(b *testing.B) {
	for idx := range recordsSearchInsert {
		for i := 0; i < b.N; i++ {
			_ = searchInsert1(recordsSearchInsert[idx].nums, recordsSearchInsert[idx].val)
		}
	}
}

func Benchmark_searchInsert2(b *testing.B) {
	for idx := range recordsSearchInsert {
		for i := 0; i < b.N; i++ {
			_ = searchInsert2(recordsSearchInsert[idx].nums, recordsSearchInsert[idx].val)
		}
	}
}
