/*Given an array of strings nums containing n unique binary strings each of length n,
return a binary string of length n that does not appear in nums.
If there are multiple answers, you may return any of them.*/

package problems

import (
	"fmt"
	"strconv"
)

func FindDifferentBinaryString(nums []string) string { // O(n)

	n := len(nums)
	bitLength := len(nums[0])      // Determine bit length from input
	totalNumbers := 1 << bitLength // 2^bitLength (all possible values) using shift to left
	// a bitwise operation , also math.Pow(2, float64(bitLength)) could have been used but it uses more
	// resources

	if n >= 1 && n <= 16 {

		allNumbers := make(map[int]bool)
		for i := 0; i < totalNumbers; i++ {
			allNumbers[i] = true
		}

		for _, numStr := range nums {
			num, err := strconv.ParseInt(numStr, 2, 64) // Convert binary string to integer
			if err != nil {
				fmt.Println("Error converting:", err)
				return ""
			}
			delete(allNumbers, int(num))
		}
		// Get the first missing number
		for missingNum := range allNumbers {
			return fmt.Sprintf("%0*b", bitLength, missingNum)
		} // this is needed because strconv.formatint() returns
		// the shortest possible representation of a binary number
		// by using sprintf i can return with a fixed lenght

	}
	return ""
}
