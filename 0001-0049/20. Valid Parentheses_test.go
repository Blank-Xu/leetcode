package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/valid-parentheses/

var valueArray = [6]struct {
	symbol byte
	value  int
}{
	{'(', 1},
	{')', -1},

	{'[', 3},
	{']', -3},

	{'{', 5},
	{'}', -5},
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 48.59% of Go online submissions for Valid Parentheses.
func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	var total int
	sl := make([]int, l/2)

	for i := 0; i < l; i++ {
		n := getValue(s[i])

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

func getValue(b byte) int {
	for idx := range valueArray {
		if b == valueArray[idx].symbol {
			return valueArray[idx].value
		}
	}

	return 0
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
	for idx := range recordsIsValid {
		for i := 0; i < b.N; i++ {
			_ = isValid(recordsIsValid[idx].input)
		}
	}
}
