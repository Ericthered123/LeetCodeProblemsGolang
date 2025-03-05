/*You are given a string title consisting of one or more words separated by a single space,
where each word consists of English letters.
Capitalize the string by changing the capitalization of each word such that:

.If the length of the word is 1 or 2 letters, change all letters to lowercase.
.Otherwise, change the first letter to uppercase and the remaining letters to lowercase.

Return the capitalized title.*/

package problems

import (
	"strings"
	"unicode"
)

func CapitalizeTitle(title string) string {
	titleField := strings.Fields(title)
	capitalizedTitle := ""
	for i := 0; i < len(titleField); i++ {
		if len(titleField[i]) == 2 || len(titleField[i]) == 1 {
			capitalizedTitle += " " + strings.ToLower(titleField[i])

		} else {
			capitalizedTitle += " " + Capitalize(titleField[i])
		}

	}
	return strings.TrimSpace(capitalizedTitle)
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes[0]) + strings.ToLower(string(runes[1:]))

}
