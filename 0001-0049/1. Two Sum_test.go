package leetcode

import (
	"github.com/Blank-Xu/leetcode-golang/utils"

	"testing"
)

// https://leetcode.com/problems/two-sum/
// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	for idx := range nums {
		for i := 0; i < idx; i++ {
			if target == nums[idx]+nums[i] {
				return []int{i, idx}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		if idx, ok := m[-v]; ok {
			return []int{idx, i}
		}

		m[v-target] = i
	}

	return nil
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsTwoSum = []struct {
	nums   []int
	target int
	output []int
}{
	{
		[]int{2, 7, 11, 15}, 9,
		[]int{0, 1},
	},
	{
		[]int{-3, 4, 3, 90}, 0,
		[]int{0, 2},
	},
	{
		[]int{-3, 1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 3, 90}, 0,
		[]int{0, 16},
	},
	{
		[]int{-3, 1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 3, 90}, 0,
		[]int{0, 30},
	},
}

func Test_twoSum(t *testing.T) {
	for _, record := range recordsTwoSum {
		if result := twoSum(record.nums, record.target); !utils.EqualSlice(result, record.output) {
			t.Errorf("twoSum test data: %v, need: %v, get: %v", record, record.output, result)
		}

		if result := twoSum2(record.nums, record.target); !utils.EqualSlice(result, record.output) {
			t.Errorf("twoSum2 test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, record := range recordsTwoSum {
			_ = twoSum(record.nums, record.target)
		}
	}
}

func Benchmark_twoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, record := range recordsTwoSum {
			_ = twoSum2(record.nums, record.target)
		}
	}
}
