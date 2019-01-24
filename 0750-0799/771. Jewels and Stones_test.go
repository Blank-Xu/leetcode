package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	var num int
	l1 := len(J)
	l2 := len(S)
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if J[i:i+1] == S[j:j+1] {
				num++
			}
		}
	}
	return num
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsNumJewelsInStones = []struct {
	j, s   string
	output int
}{
	{"aA", "aAAbbbb", 3},
	{"z", "ZZ", 0},
}

func Test_numJewelsInStones(t *testing.T) {
	for _, record := range recordsNumJewelsInStones {
		if result := numJewelsInStones(record.j, record.s); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_numJewelsInStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numJewelsInStones(recordsNumJewelsInStones[0].j, recordsNumJewelsInStones[0].s)
	}
}
