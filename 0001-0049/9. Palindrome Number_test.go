package leetcode

import (
	"strconv"
	"testing"
)

// https://leetcode.com/problems/palindrome-number/

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

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	n := x
	m := 0
	for n != 0 {
		m = m*10 + n%10
		n /= 10
	}

	return x == m
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
			t.Errorf("isPalindrome test data: %v, need: %v, get: %v", record, record.output, result)
		}

		if result := isPalindrome2(record.input); result != record.output {
			t.Errorf("isPalindrome2 test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_isPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, record := range recordsPalindrome {
			_ = isPalindrome(record.input)
		}
	}
}

func Benchmark_isPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, record := range recordsPalindrome {
			_ = isPalindrome2(record.input)
		}
	}
}
