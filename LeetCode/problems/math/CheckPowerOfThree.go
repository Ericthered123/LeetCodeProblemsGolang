/*Given an integer n, return true if it is possible to represent n
as the sum of distinct powers of three.
Otherwise, return false.

An integer y is a power of three if there exists an integer x such that y == 3x.*/

package problems

func CheckPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 2 {
			return false
		} else {
			n /= 3
		}
	}
	return true
}
