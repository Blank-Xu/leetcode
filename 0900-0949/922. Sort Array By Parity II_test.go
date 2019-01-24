package leetcode

import (
	"testing"

	"github.com/Blank-Xu/leetcode-golang/utils"
)

// https://leetcode-cn.com/problems/sort-array-by-parity-ii/

func sortArrayByParityII(A []int) []int {
	l := len(A)
	ret := make([]int, l)
	evenIdx := 0
	oddIdx := 1

	for i := 0; i < l; i++ {
		if A[i]&1 == 0 {
			ret[evenIdx] = A[i]
			evenIdx += 2
		} else {
			ret[oddIdx] = A[i]
			oddIdx += 2
		}
	}

	return ret
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsSortArrayByParityII = []struct {
	input  []int
	output []int
}{
	{[]int{4, 2, 5, 7}, []int{4, 5, 2, 7}},
}

func Test_sortArrayByParityII(t *testing.T) {
	for _, record := range recordsSortArrayByParityII {
		if result := sortArrayByParityII(record.input); !utils.EqualSlice(result, record.output) {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_sortArrayByParityII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sortArrayByParityII(recordsSortArrayByParityII[0].input)
	}
}
