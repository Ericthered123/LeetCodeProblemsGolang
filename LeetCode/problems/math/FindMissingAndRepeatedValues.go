/*You are given a 0-indexed 2D integer matrix grid of
size n * n with values in the range [1, n2].
Each integer appears exactly once except a which appears twice and b which is missing.
The task is to find the repeating and missing numbers a and b.

Return a 0-indexed integer array ans of size 2 where ans[0] equals to a and ans[1] equals to b.*/

package problems

func FindMissingAndRepeatedValues(grid [][]int) []int {
	var missingAndRepeatedValueArray [2]int
	n := len(grid)
	numberOccurrency := make(map[int]int)
	totalNumbers := n * n

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			numberOccurrency[grid[i][j]]++
		}

	}
	for j := 1; j <= totalNumbers; j++ {
		if numberOccurrency[j] == 2 {
			missingAndRepeatedValueArray[0] = j
		}
		if numberOccurrency[j] == 0 {
			missingAndRepeatedValueArray[1] = j

		}

	}
	return missingAndRepeatedValueArray[:]

}
