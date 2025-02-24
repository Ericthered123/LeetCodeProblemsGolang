/*Given a pattern and a string s, find if s follows the same pattern.

Here follow means a full match, such that there is a bijection between
 a letter in pattern and a non-empty word in s. Specifically:

.Each letter in pattern maps to exactly one unique word in s.
.Each unique word in s maps to exactly one letter in pattern.
.No two letters map to the same word, and no two words map to the same letter.
*/

package problems

import (
	"strings"
)

func WordPattern(pattern string, s string) bool {

	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false

	}
	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		if mappedWord, exists := patternToWord[char]; exists {
			if mappedWord != word {
				return false
			}

		} else {
			patternToWord[char] = word
		}

		if mappedChar, exists := wordToPattern[word]; exists {
			if mappedChar != char {
				return false
			}

		} else {
			wordToPattern[word] = char
		}
	}
	return true
}
