package strings

// https://leetcode.com/problems/split-a-string-in-balanced-strings/
func balancedStringSplit(s string) int {
	count := 0
	r := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			r = r + 1
		} else {
			r = r - 1
		}
		if r == 0 {
			count = count + 1
			r = 0
		}
	}
	return count
}
