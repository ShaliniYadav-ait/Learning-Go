package leetcode

func coinChange(coins []int, amount int) int {

	amt := make([]int, amount+1)
	diff := 0

	for i := 1; i <= amount; i++ {
		min := amount + 1
		for j := 0; j < len(coins); j++ {
			diff = i - coins[j]
			if diff < 0 {
				continue
			}

			if min > 1+amt[diff] && 1+amt[diff] != 0{
				min = 1 + amt[diff]
			}
		}

		if min != amount+1 {
			amt[i] = min
		}

		if amt[i] == 0 {
			amt[i] = -1
		}
	}
	return amt[amount]
}
