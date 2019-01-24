package leetcode

import (
	"testing"

	"github.com/Blank-Xu/leetcode-golang/utils"
)

// https://leetcode-cn.com/problems/sort-array-by-parity/

func sortArrayByParity(A []int) []int {
	l := len(A)
	even := make([]int, l)
	odd := make([]int, l)

	var evenIdx, oddIdx int

	for i := 0; i < l; i++ {
		if A[i]&1 == 0 {
			even[evenIdx] = A[i]
			evenIdx++
		} else {
			odd[oddIdx] = A[i]
			oddIdx++
		}
	}

	return append(even[:evenIdx], odd[:oddIdx]...)
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsSortArrayByParity = []struct {
	input  []int
	output []int
}{
	{[]int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
}

func Test_sortArrayByParity(t *testing.T) {
	for _, record := range recordsSortArrayByParity {
		if result := sortArrayByParity(record.input); !utils.EqualSlice(result, record.output) {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_sortArrayByParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sortArrayByParity(recordsSortArrayByParity[0].input)
	}
}
