package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/unique-morse-code-words/

func uniqueMorseRepresentations(words []string) int {
	m := make(map[string]int)
	var (
		i, j, l int
		s       string
	)
	for i = range words {
		s = ""
		l = len(words[i])
		for j = 0; j < l; j++ {
			s += wordsMap[words[i][j:j+1]]
		}
		m[s]++
	}
	return len(m)
}

var wordsMap = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsUniqueMorseRepresentations = []struct {
	input  []string
	output int
}{
	{[]string{"gin", "zen", "gig", "msg"}, 2},
}

func Test_uniqueMorseRepresentations(t *testing.T) {
	for _, record := range recordsUniqueMorseRepresentations {
		if result := uniqueMorseRepresentations(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_uniqueMorseRepresentations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uniqueMorseRepresentations(recordsUniqueMorseRepresentations[0].input)
	}
}
