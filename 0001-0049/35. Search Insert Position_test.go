package leetcode

import (
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/search-insert-position/

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

	nums = append(nums, target)
	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})

	return index
}

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

	if target > nums[l-1] {
		return l
	}

	return -1
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
	for i := 0; i < b.N; i++ {
		_ = searchInsert1(recordsSearchInsert[0].nums, recordsSearchInsert[0].val)
	}
}

func Benchmark_searchInsert2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = searchInsert2(recordsSearchInsert[0].nums, recordsSearchInsert[0].val)
	}
}
