package leetcode

import (
	"github.com/Blank-Xu/leetcode-golang/utils"

	"testing"
)

// https://leetcode-cn.com/problems/plus-one/

func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return digits
	}

	var add bool
	for i := l - 1; i >= 0; i-- {
		if add {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i]++
				add = false
			}
			continue
		}

		if i == l-1 {
			if digits[i] == 9 {
				digits[i] = 0
				add = true
			} else {
				digits[i]++
				add = false
			}
		}
	}

	if add {
		digits = append([]int{1}, digits[0:]...)
	}

	return digits
}

func plusOne2(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return digits
	}
	var (
		add   bool
		index int
	)
	for i := l; i >= 0; i-- {
		index = i - 1
		if add {
			if i == 0 {
				digits = append([]int{1}, digits[0:]...)
				return digits
			}
			if digits[index] == 9 {
				digits[index] = 0
			} else {
				digits[index]++
				add = false
			}
			continue
		}
		if i == l {
			if digits[index] == 9 {
				digits[index] = 0
				add = true
			} else {
				digits[index]++
				add = false
			}
		}
	}
	return digits
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsPlusOne = []struct {
	input  []int
	output []int
}{
	{[]int{1, 2, 3}, []int{1, 2, 4}},
	{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
}

func Test_plusOne(t *testing.T) {
	for _, record := range recordsPlusOne {
		if result := plusOne(record.input); !utils.EqualSlice(result, record.output) {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_plusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = plusOne(recordsPlusOne[0].input)
	}
}
