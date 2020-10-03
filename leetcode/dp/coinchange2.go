package leetcode

func coinChange2(coins []int, amount int) int {

total := make([][]int,len(coins)+1)

for i:= range total{
	total[i] = make([]int,amount+1)
}

total[0][0] = 1

for i := 0; i < len(coins); i++{
	for j := 0 ; j <= amount; j++{

		diff := j-coins[i]
		if diff >= 0{
			total[i+1][j] = total[i][j]+ total[i+1][diff]
			continue
		}
		total[i+1][j] = total[i][j]
	}
}
return total[len(coins)][amount]
}
