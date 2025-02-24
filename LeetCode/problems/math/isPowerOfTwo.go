/*Given an integer n, return true if it is a power of two.
Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2^x

*/

package problems

func IsPowerOfTwo(n int) bool {

	if n <= 0 {
		return false

	} else {
		return (n & (n - 1)) == 0 //if a number is a power of two then its binary representation only has 1 bit set
		// and the others are 0s
	}

	// O(1) time complexity

	/*// Compute log base 2
	logValue := math.Log2(float64(n))

	// Check if logValue is an integer
	return math.Floor(logValue) == logValue*/
}
