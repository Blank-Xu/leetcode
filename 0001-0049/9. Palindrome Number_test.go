package leetcode

import (
	"strconv"
	"testing"
)

// https://leetcode-cn.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x >= 0 && x < 10 {
		return true
	}
	b := []byte(strconv.Itoa(x))
	l := len(b)
	n := l / 2
	for i := 0; i < n; i++ {
		if b[i] != b[l-i-1] {
			return false
		}
	}
	return true
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsPalindrome = []struct {
	input  int
	output bool
}{
	{123, false},
	{-123, false},
	{0, true},
	{121, true},
	{1111, true},
}

func Test_isPalindrome(t *testing.T) {
	for _, record := range recordsPalindrome {
		if result := isPalindrome(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPalindrome(recordsReverse[0].input)
	}
}
