package leetcode

import (
	"math"
	"strconv"
	"testing"
)

// https://leetcode.com/problems/reverse-integer/
// https://leetcode-cn.com/problems/reverse-integer/

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	s := strconv.Itoa(x)
	minus := false

	if s[:1] == "-" {
		minus = true
		s = s[1:]
	}

	s = reverseString(s)
	if minus {
		s = "-" + s
	}

	x, _ = strconv.Atoi(s)

	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	return x
}

func reverseString(s string) string {
	l := len(s) - 1
	b := []byte(s)
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
// Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Reverse Integer.
func reverse2(x int) int {
	if x == 0 || x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	minus := false
	if x < 0 {
		minus = true
		x = 0 - x
	}

	n := 0
	for x != 0 {
		n = n*10 + x%10
		x /= 10
	}

	if minus {
		n = 0 - n
	}

	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
	}

	return n
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsReverse = []struct {
	input  int
	output int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
}

func Test_reverse(t *testing.T) {
	for _, record := range recordsReverse {
		if result := reverse(record.input); result != record.output {
			t.Errorf("reverse test data: %v, need: %v, get: %v", record, record.output, result)
		}

		if result := reverse2(record.input); result != record.output {
			t.Errorf("reverse2 test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, record := range recordsReverse {
			_ = reverse(record.input)
		}
	}
}

func Benchmark_reverse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, record := range recordsReverse {
			_ = reverse2(record.input)
		}
	}
}
