/*Given two strings s and t, return true if t is an
anagram
 of s, and false otherwise.*/

package problems

import (
	"strings"
)

func IsAnagram(s string, t string) bool {
	s, t = strings.ToLower(s), strings.ToLower(t)
	s = removingNonAlpha(s)

	t = removingNonAlpha(t)
	if len(s) != len(t) {
		return false
	} else {

		count := make(map[rune]int)

		for _, char := range s {
			count[char]++
		}

		for _, char := range t {
			count[char]--
			if count[char] < 0 {
				return false
			}
		}
		return true
	} //O(n)

}
