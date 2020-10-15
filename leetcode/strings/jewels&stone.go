package strings

// https://leetcode.com/problems/jewels-and-stones/
func numJewelsInStones(J string, S string) int {
	
	count := 0
	for _, jewel := range J{
		for _, stone := range S{
			if jewel == stone{
				count++
			}
		}
	}
	return count

}