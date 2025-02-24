/*Given a positive integer num, return true if num is a perfect square or false otherwise.

A perfect square is an integer that is the square of an integer.
In other words, it is the product of some integer with itself.

You must not use any built-in library function, such as sqrt.*/

package problems

import "math"

func IsPerfectSquare(num int) bool {
	if num >= 1 && num <= math.MaxInt32 {

		left, right := 1, num

		for left <= right {
			mid := left + (right-left)/2
			if mid*mid == num {
				return true
			} else if mid*mid < num {
				left = mid + 1
			} else {
				right = mid - 1 // otherwise move right to left
			}

		}
		return false
	}
	return false
}
