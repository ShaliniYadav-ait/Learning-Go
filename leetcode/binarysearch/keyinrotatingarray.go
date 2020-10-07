package binarysearch

func findKeyInRotatingArray(vals []int, key int) int {

	start := 0
	end := len(vals) - 1
	rotation := findRotationPoint(vals)

	for start <= end {

		mid := start + (end-start)/2

		realMid := (mid + rotation) % len(vals)

		if vals[realMid] == key {
			return realMid
		}

		if vals[realMid] > key {
			end = mid - 1
			continue
		}

		start = mid + 1

	}

	return -1
}
