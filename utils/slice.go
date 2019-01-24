package utils

// https://github.com/kenshinsyrup/AllGolangDemo/tree/master/CompareSlice

// EqualSlice  check slice value
func EqualSlice(s1, s2 []int) bool {
	if (s1 == nil) != (s2 == nil) {
		return false
	}

	l := len(s1)
	if l != len(s2) {
		return false
	}

	s2 = s2[:l]
	for idx := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}

	return true
}
