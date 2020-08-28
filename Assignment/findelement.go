package assignment

func numInArray(list []int, n int) bool {

	for i := range list {
		if list[i] == n {
			return true
		}
	}

	return false
}
