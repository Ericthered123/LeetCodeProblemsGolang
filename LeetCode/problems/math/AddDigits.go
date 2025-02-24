/*Given an integer num, repeatedly add all its digits until the result has only one digit,
and return it.

*/

package problems

func AddDigits(num int) int {

	digit := 0
	result := 0
	for num/10 != 0 {
		for num != 0 {
			digit = num % 10
			num = num / 10
			result += digit
		}

	}
	return result

}
