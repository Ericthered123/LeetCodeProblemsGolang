/*Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1],
then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).*/

package problems

import "math"

func Reverse(x int32) int32 {

	var reversed int32 = 0

	for x != 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x = x / 10

	}

	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	} else {
		return reversed
	}

}
