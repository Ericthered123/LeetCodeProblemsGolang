/*Given two integers that represent binaries 1s and 05 a and b return the sum */
package problems

import (
	"math/bits"
	"strconv"
)

func BinarySumInt(a, b uint) string {
	var carry uint

	for b != 0 {
		a, carry = bits.Add(a, b, 0)
		b = carry

	}
	var result uint = uint(a)
	return strconv.FormatUint(uint64(result), 2)

}
