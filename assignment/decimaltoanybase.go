package assignment

func decimalToAnyBase(number int, base int,digits []string) string {
var conversion string

if number == 0{
	return digits[0]
}
for number > 0 {

	remainder := number%base
	number = number / base

	conversion = digits[remainder] + conversion 

}
return conversion
}