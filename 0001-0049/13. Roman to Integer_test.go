package leetcode

import (
	"testing"
)

// https://leetcode.com/problems/roman-to-integer/
// https://leetcode-cn.com/problems/roman-to-integer/

var (
	romanArray = [7]struct {
		symbol byte
		num    int
	}{
		{'I', 1},
		{'V', 5},
		{'X', 10},
		{'L', 50},
		{'C', 100},
		{'D', 500},
		{'M', 1000},
	}

	romanRules = [3]struct {
		symbol byte
		r1, r2 byte
		v1, v2 int
	}{
		{'I', 'V', 'X', 4, 9},
		{'X', 'L', 'C', 40, 90},
		{'C', 'D', 'M', 400, 900},
	}
)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Roman to Integer.
// Memory Usage: 3.1 MB, less than 79.06% of Go online submissions for Roman to Integer.
func romanToInt(s string) int {
	var v int
	l := len(s)

	switch l {
	case 1:
		return romanGetValue(s[0])
	case 2:
		v = romanCheckRule(s[0], s[1])
		if v > 0 {
			return v
		}
	}

	b := []byte(s)
	n := l - 1

	var (
		skip  bool
		value int
	)

	for i := 0; i < n; i++ {
		if skip {
			skip = false
			continue
		}

		v = romanCheckRule(b[i], b[i+1])

		if v > 0 {
			skip = true
			value += v
		} else {
			value += romanGetValue(b[i])
		}
	}

	if !skip {
		value += romanGetValue(b[n])
	}

	return value
}

func romanCheckRule(l, r byte) int {
	for idx := range romanRules {
		if l == romanRules[idx].symbol {
			if r == romanRules[idx].r2 {
				return romanRules[idx].v2
			} else if r == romanRules[idx].r1 {
				return romanRules[idx].v1
			}
		}
	}

	return 0
}

func romanGetValue(b byte) int {
	for idx := range romanArray {
		if b == romanArray[idx].symbol {
			return romanArray[idx].num
		}
	}

	return 0
}

func romanToInt2(s string) int {
	l := len(s)
	if l == 1 {
		return romanGetValue(s[0])
	}

	var v, value int
	n := l - 1

	for i := 0; i < l; i++ {
		v = romanGetValue(s[i])
		if i == n || romanGetValue(s[i+1]) <= v {
			value += v
		} else {
			value -= v
		}
	}

	return value
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsRomanToInt = []struct {
	input  string
	output int
}{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"MDCXCV", 1695},
	{"MI", 1001},
}

func Test_romanToInt(t *testing.T) {
	for _, record := range recordsRomanToInt {
		if result := romanToInt(record.input); result != record.output {
			t.Errorf("romanToInt test data: %v, need: %v, get: %v", record, record.output, result)
		}

		if result := romanToInt2(record.input); result != record.output {
			t.Errorf("romanToInt2 test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_romanToInt(b *testing.B) {
	for _, record := range recordsRomanToInt {
		for i := 0; i < b.N; i++ {
			_ = romanToInt(record.input)
		}
	}
}

func Benchmark_romanToInt2(b *testing.B) {
	for _, record := range recordsRomanToInt {
		for i := 0; i < b.N; i++ {
			_ = romanToInt2(record.input)
		}
	}
}
