/*Given an integer x, return true if x is a
palindrome , and false otherwise.*/

package problems

func IsPalindrome(x int) bool {

	if x < 0 {
		return false

	}

	reversed := 0
	temp2 := x

	for temp2 != 0 {
		temp := temp2 % 10
		temp2 = temp2 / 10
		reversed = reversed*10 + temp

	}

	if reversed == x {
		return true

	}
	return false

}
