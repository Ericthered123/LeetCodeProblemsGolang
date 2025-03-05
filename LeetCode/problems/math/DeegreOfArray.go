/*Given a non-empty array of non-negative integers nums,
the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums,
that has the same degree as nums.*/

package problems

func FindShortestSubArray(nums []int) int {
	freq := make(map[int]int)
	firstOccurrence := make(map[int]int)
	lastOccurrence := make(map[int]int)

	for i, num := range nums {
		if _, exists := firstOccurrence[num]; !exists {
			firstOccurrence[num] = i
		}
		lastOccurrence[num] = i
		freq[num]++
	}

	maxDegree := 0
	for _, count := range freq {
		if count > maxDegree {
			maxDegree = count
		}

	}

	minLength := len(nums)
	for num, count := range freq {
		if count == maxDegree { // Only check elements that contribute to the degree
			length := lastOccurrence[num] - firstOccurrence[num] + 1
			if length < minLength {
				minLength = length
			}
		}
	}

	return minLength
}
