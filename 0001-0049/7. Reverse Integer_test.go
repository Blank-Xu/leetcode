package leetcode

import (
	"math"
	"strconv"
	"testing"
)

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
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverse(recordsReverse[0].input)
	}
}
