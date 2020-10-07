package binarysearch

func findRotationPointInDuplicate(vals []int) int {

	start := 0
	end := len(vals) -1

	if vals[start] < vals[end] {
		return 0
	}

	for start < end {
		mid := start + (end-start)/2

		if vals[mid] > vals[end] {
			start = mid + 1
		}
		
		if vals[mid] > vals[start]{
			end = mid 
		}

		if vals[mid] == vals[start] && vals[mid] == vals[end]{
			valsfirsthalf := findRotationPointInDuplicate(vals[:vals[mid]])
			valsotherhalf := findRotationPointInDuplicate(vals[vals[mid]:])

			if valsfirsthalf == 0{
				return valsotherhalf
			}
			return valsfirsthalf
		}
	}
	return start
}
