/*Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.

*/

package problems

import (
	"sort"
)

func ContainsDuplicate(nums []int) bool {

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true

		}
	}
	return false

}
