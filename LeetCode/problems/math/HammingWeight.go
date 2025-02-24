/*Given a positive integer n, write a function that returns the number of
set bits
 in its binary representation (also known as the Hamming weight).*/

package problems

import "math"

func HammingWeight(n int) int {
	if n >= 1 && n <= math.MaxInt32 {
		count := 0
		for n > 0 {
			n &= (n - 1) // clearing the ones in the number, each time a one is cleaned we count it till we
			// reach zero in binary
			count++

		}
		return count

	} //this is known as the Brian Kernighan algorithm and it works with a O(k) time  where k is the numbers of ones
	return 0
}
