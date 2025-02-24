/*Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

*/

package problems

func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs[1:] {
		for len(prefix) > len(str) || len(str) > 0 && str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]

		}
		if prefix == "" {
			return ""

		}
	}
	return prefix

}
