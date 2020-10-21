package array

//https://leetcode.com/problems/can-place-flowers/
func canPlaceFlowers(flowerbed []int, n int) bool {

	if n == 0 {
		return true
	}

	count := 0
	for i := 0; i < len(flowerbed); i++ {
		Previous, Next := 0, 0
		if i > 0 && i < len(flowerbed)-1 {
			Previous = i - 1
			Next = i + 1
		}

		if flowerbed[Previous]+flowerbed[i]+flowerbed[Next] == 0 {
			i++
			count++
		}

		if count == n {
			return true
		}

	}
	return false

}

/*
	if len(flowerbed) == 0 {
		return false
	}

	if n == 0 {
		return true
	}

	index := 0
	indexflag := false

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && indexflag == false {
			index = i
			indexflag = true
			continue
		}

		if flowerbed[i] == 1 && (indexflag == false || index == 0) {
			count := (i - index) / 2
			if n > count {
				n = n - count
				indexflag = false
				continue
			}
			return true
		}

		if flowerbed[i] == 1 && indexflag == true {
			count := (i - index - 1) / 2
			if n > count {
				n = n - count
				indexflag = false
				continue
			}
			return true
		}
	}

	if len(flowerbed) > 1 {
		if flowerbed[len(flowerbed)-1] == 0 && indexflag != false {
			if index == 0 {
				count := (len(flowerbed) - index + 1) / 2
				if (n - count) <= 0 {
					return true
				}
			}
			count := (len(flowerbed) - index) / 2
			if (n - count) <= 0 {
				return true
			}
		}
	} else {
		if indexflag == true && n == 1 {
			return true
		}
	}

	return false
}
*/
