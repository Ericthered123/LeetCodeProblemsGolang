/*Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order.
*/

package problems

func TwoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}

	} //It iterates n times in the first for, and n/2 times in the second for so the time complexity is
	// n* n/2 = n^2 ---> O(n^2) to better this time complexity we can optimize this function by using
	// maps.

	return []int{}

}
