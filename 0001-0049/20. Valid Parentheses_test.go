package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/valid-parentheses/

var valueMap = map[string]int{
	"(": 1,
	")": -1,

	"[": 3,
	"]": -3,

	"{": 5,
	"}": -5,
}

func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	var total int
	sl := make([]int, l/2)
	for i := 0; i < l; i++ {
		n := valueMap[string(s[i])]
		if n < 0 {
			if i == 0 {
				return false
			}
			j := len(sl) - 1
			if n+sl[j] != 0 {
				return false
			}
			sl = sl[0:j]
		} else {
			sl = append(sl, n)
		}
		total += n
	}
	return total == 0
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsIsValid = []struct {
	input  string
	output bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
}

func Test_isValid(t *testing.T) {
	for _, record := range recordsIsValid {
		if result := isValid(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isValid(recordsIsValid[0].input)
	}
}
