/*A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
and removing all non-alphanumeric characters,  it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

*/

package problems

import (
	"strings"
	"unicode"
)

func IsPalindromeString(s string) bool {

	s = strings.ToLower(s)
	s = removingNonAlpha(s)
	reversedString := reverseString(s)
	for i := 0; i < len(s); i++ {
		if s[i] != reversedString[i] {
			return false
		}
	}

	return true

}

/*func joinString(s string) string {
	fields := strings.Fields(s)
	joinedString := ""
	for i := 0; i < len(fields); i++ {
		joinedString += fields[i]
	}

	return joinedString

}*/

func removingNonAlpha(s string) string {

	var builder strings.Builder

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(char)
		}
	}

	return builder.String()

}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
