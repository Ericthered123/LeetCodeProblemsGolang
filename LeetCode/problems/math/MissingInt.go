/*Given an array nums containing n distinct numbers in the range [0, n],
return the only number in the range that is missing from the array.*/

package problems

func MissingNumber(nums []int) int {
	actualXOR := 0
	expectedXOR := 0
	for i := 0; i < len(nums)+1; i++ {
		expectedXOR ^= i
	}
	for i := 0; i < len(nums); i++ {
		actualXOR = nums[i] ^ actualXOR
	}

	missingNum := expectedXOR ^ actualXOR
	return missingNum

}
