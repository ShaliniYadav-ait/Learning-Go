package others

func replaceAndDelete(s []rune, k int) []rune {

	countA,countB := 0,0
	w := 0
	for i := 0; i < k; i++ {
		if s[i] == 'a' {
			countA++
		}

		if s[i] != 'b' {
			s[w] = s[i]
			w++
			continue
		}
		countB++
	}

	w = k - countB + countA - 1
	for i := k - countB - 1 ; i >= 0; i-- {
		if s[i] != 'a' {
			s[w] = s[i]
			w--
			continue
		}

		s[w] = 'd'
		s[w-1] = 'd'
		w = w - 2

	}

	return s[:k+countA-countB]
}
