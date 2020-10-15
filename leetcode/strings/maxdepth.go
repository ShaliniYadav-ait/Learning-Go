package strings

// https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
func maxDepth(s string) int {

	count, max := 0, 0
	for _, val := range s {

		if val == '(' {
			count++
		}
		if val == ')' {
			count--
		}

		if max < count{
			max = count
		}
	}
return max

}
