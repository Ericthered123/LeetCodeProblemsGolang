/*Given an integer array nums and an integer k,
return true if there are two distinct indices i and j
in the array such that nums[i] == nums[j] and abs(i - j) <= k.	*/

package problems

func ContainsNearbyDuplicate(nums []int, k int) bool {
	if ContainsDuplicate(nums) {
		indexMap := make(map[int]int)

		for i, num := range nums {
			if prevIndex, exists := indexMap[num]; exists && i-prevIndex <= k {
				return true
			}
			indexMap[num] = i
		}
		return false

	}
	return false
}
