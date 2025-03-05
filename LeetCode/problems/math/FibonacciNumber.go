/*The Fibonacci numbers, commonly denoted F(n) form a sequence,
called the Fibonacci sequence, such that each number is the sum of the two preceding ones,
starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).*/

package problems

/*func Fib(n int) int {
if n == 0 {
	return 0

}
if n == 1 {
	return 1
}
return Fib(n-1) + Fib(n-2)
}
*/
//Complexity of O(2^n) because of using simple recursion
// this could be upgraded to O(n) by using cache or memoization or just by doing dynamic progamming

var memo = make(map[int]int) //cache to avoid repetitive calculation

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if val, exists := memo[n]; exists { // if fib(n) has already been calculated, then we return it
		return val
	}
	memo[n] = Fib(n-1) + Fib(n-2) // Storing the result
	return memo[n]
}
