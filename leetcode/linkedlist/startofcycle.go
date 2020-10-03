package linkedlist

import (
	"math"
)

func startOfCycle(vals []*ListNode) *ListNode {

	lengthVals := [2]int{}
	hasCycle := [2]bool{}
	checkIntersection := 0
	min, max := math.MaxInt32, math.MinInt32
	index,minIndex := -1,-1

	for i := 0; i < len(vals); i++ {
		lengthVals[i], hasCycle[i] = lengthOfList(vals[i])
	}

	for i := 0; i < len(hasCycle); i++ {
		if hasCycle[i] == true {
			checkIntersection++
		}
	}

	if checkIntersection%2 != 0{
		return nil
	}

	for i := 0; i < len(lengthVals); i++ {
		if min > lengthVals[i] {
			min = lengthVals[i]
			minIndex = i
		}

		if max <= lengthVals[i]{
			max = lengthVals[i]
			index = i
		}
	}

	diff := max - min

	var move *ListNode
	move = vals[index]

for diff != 0{
	move = move.Next
	diff--
}

var other *ListNode
other = vals[minIndex]

for other != move{
	other = other.Next
	move = move.Next
}

return move

}
