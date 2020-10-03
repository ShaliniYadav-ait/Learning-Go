package binarysearch

func findCeil(vals []int, key int) int {

	start := 0
	end := len(vals) - 1
	immediateNext := -1

	for start <= end {
		mid := start + (end-start)/2

		if vals[mid] == key {
			return mid
		}

		if vals[mid] < key {
			start = mid + 1
			continue
		}

		immediateNext = mid
		end = mid - 1
	}
	return immediateNext
}
