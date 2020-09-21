package array

func maxProfit(prices []int) int {

	min, profit := prices[0], 0

	for _, val := range prices {
	
		if val < min{
			min = val
		}

		if val-min > profit{
			profit = val - min
		}
	}

	return profit
}
