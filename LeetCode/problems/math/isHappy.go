/*Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.*/

package problems

func IsHappy(n int) bool {

	seen := make(map[int]bool)

	for n != 1 {
		if seen[n] {
			return false
		}

		for n > 0 {
			sumaDig += (n % 10) * (n % 10)
			n /= 10

		}
		n = sumaDig
	}
	return true
}
