package strings

//https://leetcode.com/problems/shuffle-string/
func restoreString(s string, indices []int) string {

	result := make([]string, len(indices))
	for i, val := range s {
		result[indices[i]] = string(val)
	}
	
	s = ""
	for i := 0; i < len(result); i++ {
		s = s + result[i]
	}

	return s
}
