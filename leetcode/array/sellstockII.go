package array

import (
	"fmt"
)

func maxProfit2(prices []int) int {

	profit := make([][]int, 2)
	min := prices[0]

	fmt.Println(len(prices))
	for i := range profit {
		profit[i] = make([]int, len(prices))
	}

	for i := 1; i < len(prices); i++ {
		profit[0][i] = profit[0][i-1]
		if prices[i]-min > profit[0][i-1] {
			profit[0][i] = prices[i] - min
			continue
		}
		
        if prices[i]< min{
			min = prices[i]
		  }
	}

	for i := 1; i < len(prices); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if prices[j] < prices[i] {
				currProfit := prices[i] - prices[j]
				if j != 0 {
					currProfit = currProfit + profit[0][j-1]
				}
				if currProfit > max {
					max = currProfit
				}
			}
		}
		if max > profit[1][i-1] {
			profit[1][i] = max
			continue
		}
		profit[1][i] = profit[1][i-1]
	}
	return profit[1][len(prices)-1]
}
