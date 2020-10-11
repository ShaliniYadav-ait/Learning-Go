package glider

func findElementInNearlySortedArray(vals []int, key int) int {

	start := 0
	end := len(vals) - 1

	for {
		mid := start + (end-start)/2

		if vals[mid] == key {
			return mid
		}

		if mid > start && vals[mid-1] == key {
			return mid - 1
		}

		if mid < end && vals[mid+1] == key {
			return mid + 1
		}

		if key > vals[mid+1] {
			start = mid + 2
		}

		if key < vals[mid-1] {
			end = mid - 1
		}
	}

}
