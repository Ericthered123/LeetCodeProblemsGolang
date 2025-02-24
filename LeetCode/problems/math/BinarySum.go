/*Given two binary strings a and b, return their sum as a binary string.*/

package problems

import (
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
func AddBinary(a string, b string) string {

	i, j := len(a)-1, len(b)-1
	carry := 0
	var result strings.Builder

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i > 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result.WriteByte(byte(sum%2) + '0')
		carry = sum / 2
	}
	binarySum := result.String()
	return reverseString(binarySum)

}
