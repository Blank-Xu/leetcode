package leetcode

import (
	"testing"
)

// https://leetcode-cn.com/problems/robot-return-to-origin/

func judgeCircle(moves string) bool {
	var top, side int
	l := len(moves)
	for i := 0; i < l; i++ {
		switch moves[i : i+1] {
		case "U":
			top++
		case "D":
			top--
		case "L":
			side++
		case "R":
			side--
		}
	}
	return top == 0 && side == 0
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsJudgeCircle = []struct {
	input  string
	output bool
}{
	{"UD", true},
	{"LL", false},
}

func Test_judgeCircle(t *testing.T) {
	for _, record := range recordsJudgeCircle {
		if result := judgeCircle(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_judgeCircle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = judgeCircle(recordsJudgeCircle[0].input)
	}
}
