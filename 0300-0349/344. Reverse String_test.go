package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/reverse-string/

func reverseString(s string) string {
	l := len(s) - 1
	b := []byte(s)
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func reverseString2(s string) string {
	l := len(s)
	var ret string
	for i := l; i > 0; i-- {
		ret += s[i-1 : i]
	}
	return ret
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsReverseString = []struct {
	input  string
	output string
}{
	{"hello", "olleh"},
	{"Hannah", "hannaH"},
}

func Test_reverseString(t *testing.T) {
	for _, record := range recordsReverseString {
		if result := reverseString(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_reverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseString(recordsReverseString[0].input)
	}
}
