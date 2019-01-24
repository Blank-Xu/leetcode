package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/to-lower-case/

func toLowerCase(str string) string {
	l := len(str)
	var (
		ok     bool
		s, ret string
	)
	for i := 0; i < l; i++ {
		if s, ok = letterMap[str[i:i+1]]; ok {
			ret += s
		} else {
			ret += str[i : i+1]
		}
	}
	return ret
}

var letterMap = map[string]string{
	"A": "a",
	"B": "b",
	"C": "c",
	"D": "d",
	"E": "e",
	"F": "f",
	"G": "g",
	"H": "h",
	"I": "i",
	"J": "j",
	"K": "k",
	"L": "l",
	"M": "m",
	"N": "n",
	"O": "o",
	"P": "p",
	"Q": "q",
	"R": "r",
	"S": "s",
	"T": "t",
	"U": "u",
	"V": "v",
	"W": "w",
	"X": "x",
	"Y": "y",
	"Z": "z",
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsToLowerCase = []struct {
	input  string
	output string
}{
	{"Hello", "hello"},
	{"here", "here"},
	{"LOVELY", "lovely"},
}

func Test_toLowerCase(t *testing.T) {
	for _, record := range recordsToLowerCase {
		if result := toLowerCase(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_toLowerCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toLowerCase(recordsToLowerCase[0].input)
	}
}
