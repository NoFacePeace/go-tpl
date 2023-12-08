package digit

func GetDigit(num int) []int {
	ret := []int{}
	for num != 0 {
		mod := num % 10
		num /= 10
		ret = append(ret, mod)
	}
	return ret
}
