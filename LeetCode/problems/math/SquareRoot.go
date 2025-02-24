/*Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.*/

package problems

import "math"

func MySqrt(x int) int {

	if x >= 0 && x <= (math.MaxInt32) { //i ll solve it by using binary search since we are working with integers

		if x == 0 {
			return 0

		}
		left, right := 1, x

		var ans int

		for left <= right {
			mid := left + (right-left)/2 // this would be the middle point in the set
			if mid*mid == x {            // this would be the desired and best outcome for time complexity
				return mid
			} else if mid*mid < x { // if this happens then we should have to move left to the right
				ans = mid
				left = mid + 1
			} else {
				right = mid - 1 // otherwise move right to left
			}
		}

		return ans // here i return the last valid mid rounded down

	}
	return 0
}
