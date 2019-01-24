package leetcode

import (
	"strings"
	"testing"
)

// https://leetcode-cn.com/problems/unique-email-addresses/

func numUniqueEmails(emails []string) int {
	if len(emails) == 0 {
		return 0
	}
	var formatEmail = func(email string) string {
		idx1 := strings.Index(email, "+")
		idx2 := strings.Index(email, "@")
		var s string
		if idx1 > 0 {
			s = email[0:idx1]
		} else {
			s = email[0:idx2]
		}
		s = strings.Replace(s, ".", "", -1)
		return s + email[idx2:]
	}
	m := make(map[string]bool)
	for i := range emails {
		m[formatEmail(emails[i])] = true
	}
	return len(m)
}

// ------------------------------------------------
// ------------test and benchmark------------------
// ------------------------------------------------

var recordsNumUniqueEmails = []struct {
	input  []string
	output int
}{
	{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
}

func Test_numUniqueEmails(t *testing.T) {
	for _, record := range recordsNumUniqueEmails {
		if result := numUniqueEmails(record.input); result != record.output {
			t.Errorf("test data: %v, need: %v, get: %v", record, record.output, result)
		}
	}
}

func Benchmark_numUniqueEmails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numUniqueEmails(recordsNumUniqueEmails[0].input)
	}
}
