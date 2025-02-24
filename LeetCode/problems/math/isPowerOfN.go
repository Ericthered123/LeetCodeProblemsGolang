package problems

func IsPowerOfN(num int, exp int) bool { // Only for positive integers numbers including zero

	if num <= 0 || exp < 0 {
		return false
	}
	if exp == 0 && num == 1 {
		return true
	} else if exp == 0 && num != 1 {
		return false
	}
	for num%exp == 0 {
		num /= exp
	}

	return num == 1

}
