package others

//https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
func subtractProductAndSum(n int) int {
 
	sum , prod := 0,0
	for n != 0{
		mod := n%10
		n = n/10

		sum = sum + mod
		prod = prod * mod

	}

	return prod - sum
}