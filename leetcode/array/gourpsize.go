package array

import (
	"fmt"
	"strconv"
)

//https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
func groupThePeople(groupSizes []int) [][]int {

	result := [][]int{}
	group := make(map[int]string)
	max := 0

	for i := 0; i < len(groupSizes); i++ {

		if groupSizes[i] > max {
			max = groupSizes[i]
		}

		_, ok := group[groupSizes[i]]
		if ok {
			group[groupSizes[i]] = group[groupSizes[i]] + "|" + fmt.Sprint(i)
			continue
		}

		group[groupSizes[i]] = fmt.Sprint(i)

	}

	for i := 0; i <= max; i++ {
		_, ok := group[i]
		if !ok {
			continue
		}
		count := 0
		var pos string
		res1 := []int{}
		for _, index := range group[i] {

			if index != '|' {
				pos = pos + string(index)
				continue
			}

				count++
			res, _ := strconv.Atoi(pos)
			res1 = append(res1, res)
			pos = ""

			if count == i {
				result = append(result, res1)
				count = 0
				res1 = nil
			}

		}
		if (count != 0 && res1 != nil) || pos != "" {
				res, _ := strconv.Atoi(pos)
				res1 = append(res1, res)

					result = append(result, res1)
					count = 0
					res1 = nil
		}

	}

	return result
}
