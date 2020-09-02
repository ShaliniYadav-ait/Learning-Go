package assignment

func anyBaseToDecimal(s string, base int, digits []string) int {

	val := 0
	valindex := 0
	power := len(s) - 1

	for i:= 0;i<len(s);i++ {

		for j, conv := range digits {
			if conv == string(s[i]) {
				valindex = j
				break
			}
		}

		val = val + valindex*(recursivePower(base, power))
		power--
	}
	return val
}

func recursivePower(base int, exponent int) int {
	if exponent != 0 {
		return (base * recursivePower(base, exponent-1))
	}
	return 1
}
