package strings

//https://leetcode.com/problems/valid-palindrome-ii/
func validPalindrome(s string) bool {
	var palindrome bool
	count := 1
	r := []rune(s)

	palindrome = checkPalindrome(r, count)

	return palindrome
}

func checkPalindrome(r []rune, count int) bool {

	size := len(r)
	last := size - 1

	if size <= 1{
		return true
	}

	if r[0] == r[last] {
		return checkPalindrome(r[1:last], count)
	} else if count > 0 {
		return checkPalindrome(r[0:last], count-1) || checkPalindrome(r[1:last+1], count-1)
	}
	return false

}
