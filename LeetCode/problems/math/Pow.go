package problems

func MyPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / MyPow(x, -n)
	}
	if n%2 == 0 {
		half := MyPow(x, n/2) // exponentiation by squaring
		return half * half
	}
	return x * MyPow(x, n-1)

}
