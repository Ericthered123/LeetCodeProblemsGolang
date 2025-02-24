/*Given a roman numeral, convert it to an integer.*/

package problems

import "strings"

func verifyRoman(s string) bool {
	romanLetters := "IVXLCDM"
	if !strings.ContainsAny(s, romanLetters) {
		return false
	}
	return true
}

func RomanToInt(s string) int {

	if len(s) >= 1 && len(s) <= 15 && verifyRoman(s) {

		var romanConvertedToInt int = 0
		s = s + " "
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case 'I':
				if s[i+1] == 'V' || s[i+1] == 'X' {
					romanConvertedToInt -= 1

				} else {
					romanConvertedToInt += 1
				}

			case 'V':
				romanConvertedToInt += 5
			case 'X':
				if s[i+1] == 'L' || s[i+1] == 'C' {
					romanConvertedToInt -= 10
				} else {
					romanConvertedToInt += 10
				}

			case 'L':
				romanConvertedToInt += 50
			case 'C':
				if s[i+1] == 'D' || s[i+1] == 'M' {
					romanConvertedToInt -= 100
				} else {
					romanConvertedToInt += 100
				}
			case 'D':
				romanConvertedToInt += 500
			case 'M':
				romanConvertedToInt += 1000
			default:
				romanConvertedToInt += 0
			}

		}

		return romanConvertedToInt
	}
	return 0
}
