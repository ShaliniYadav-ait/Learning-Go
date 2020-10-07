package binarysearch

func peakIndexInMountainArray(arr []int) int {
	
	start := 0
	end := len(arr) - 1

	for start <= end{

		mid := start + (end-start)/2

		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1]{
			return mid
		}

		if arr[mid] > arr[mid+1] && arr[mid] < arr[mid-1]{
			end = mid - 1
			continue
		}

		start = mid + 1
	}

	return start
}