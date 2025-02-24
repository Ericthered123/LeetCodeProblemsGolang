/*Given a non-empty array of integers nums, every element appears twice except for one.
Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.*/

package problems

func SingleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {

		result = nums[i] ^ result

	}
	return result

}
