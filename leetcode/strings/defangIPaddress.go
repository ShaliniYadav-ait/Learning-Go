package leetcode

func defangIPaddr(address string) string {

	var addr string
	for _, i := range address {
		if string(i) == "." {
			addr = addr + "[" + string(i) + "]"
			continue
		}

		addr = addr + string(i)
	}

	return addr
}
