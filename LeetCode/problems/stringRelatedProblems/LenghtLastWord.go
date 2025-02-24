/*Given a string s consisting of words and spaces,
 return the length of the last word in the string.

A word is a maximal  substring consisting of non-space characters only.*/

package problems

import (
	"strings"
)

func LengthOfLastWord(s string) int {

	if strings.Contains(s, " ") {
		s = strings.TrimSpace(s)
		n := strings.Count(s, " ")
		for i := 0; i < n; i++ {
			s = s[strings.Index(s, " ")+1:]
		}

		return len(s)

	} else {
		return len(s)
	}

}
