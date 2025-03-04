/*Given an array of points on the X-Y plane points where points[i] = [xi, yi],
return the area of the largest triangle that can be formed by any three different points.
Answers within 10^(-5) of the actual answer will be accepted.

*/

package problems

import "math"

func LargestTriangleArea(points [][]int) float64 {
	n := len(points)
	maxArea := 0.0

	// Iterate over all combinations of three points
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				// Compute the area using the shoelace formula
				x1, y1 := points[i][0], points[i][1]
				x2, y2 := points[j][0], points[j][1]
				x3, y3 := points[k][0], points[k][1]

				area := 0.5 * math.Abs(float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)))
				maxArea = math.Max(maxArea, area)
			}
		}
	} //O(n^3) since n<=50 the worst case is O(50^3)which is manageable, this complexity could be reduced
	// to O(n Log(n))

	return maxArea
}
