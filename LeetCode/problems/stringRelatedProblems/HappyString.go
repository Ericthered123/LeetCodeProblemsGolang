/*A happy string is a string that:

consists only of letters of the set ['a', 'b', 'c'].
s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is 1-indexed).
For example, strings "abc", "ac", "b" and "abcbabcbcb" are
all happy strings and strings "aa", "baa" and "ababbc" are not happy strings.

Given two integers n and k,
consider a list of all happy strings of length n sorted in lexicographical order.

Return the kth string of this list or return an empty string if there are less
than k happy strings of length n.*/

package problems

func GetHappyString(n int, k int) string {

	var result string //Stores the kth string when/if found

	letters := []rune{'a', 'b', 'c'}
	count := 0 //counting for how many strings we have built

	var backtrack func(current string) // current is the string built till now
	backtrack = func(current string) { // this recursive function add letters one by one, following the condition
		// of a happy string

		if len(current) == n {
			count++

			if count == k {
				result = current
			}
			return
		}

		for _, ch := range letters { // here the combinations are generated
			if len(current) == 0 || rune(current[len(current)-1]) != ch {
				backtrack(current + string(ch))
				if result != "" {
					return
				}
			}
		}
	}
	backtrack("")
	return result

} // worst case scenario this is O(3^n)
