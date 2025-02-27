/*Given an integer n, return true if it is a power of three.
Otherwise, return false.

An integer n is a power of three, if there exists an integer x such that n == 3x.*/

package problems

func IsPowerOfThree(n int) bool {

	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3

	}

	return n == 1
}
