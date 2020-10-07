package array

func numWaterBottles(numBottles int, numExchange int) int {

	exchange, remaining, res := 0, numBottles, 0
	tot := exchange + remaining
	for tot   >= numExchange {
		
		remaining = tot % numExchange
		exchange := tot / numExchange

		res = res + exchange

		tot = exchange + remaining
	}
	return numBottles + res
}
