package leetcode

import (
	"github.com/Blank-Xu/leetcode-golang/utils"

	"testing"
)

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if target == nums[i]+nums[j] {
				return []int{j, i}
			}
		}
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
}

func Test_twoSum(t *testing.T) {
	for _, record := range recordsTwoSum {
		if result := twoSum(record.nums, record.target); !utils.EqualSlice(result, record.output) {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = twoSum(recordsTwoSum[0].nums, recordsTwoSum[0].target)
	}
}
