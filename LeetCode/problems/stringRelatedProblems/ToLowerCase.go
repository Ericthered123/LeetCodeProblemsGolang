/*Given a string s, return the string after
replacing every uppercase letter with the same lowercase letter.*/

package problems

func ToLowerCase(s string) string {

	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] <= 90 && runes[i] >= 65 {
			runes[i] = runes[i] + 32
		}
	}
	lowerCaseString := string(runes)
	return lowerCaseString

}
