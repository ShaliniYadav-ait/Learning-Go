package assignment

func reverseString(input string) string {

	var reverse string
	for _, i := range input {
		reverse = string(i) + reverse
	}
	return reverse
}
