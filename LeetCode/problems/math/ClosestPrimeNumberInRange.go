/*Given two positive integers left and right, find the two integers num1 and num2 such that:

	left <= num1 < num2 <= right .

Both num1 and num2 are prime numbers.
num2 - num1 is the minimum amongst all other pairs satisfying the above conditions.
.Return the positive integer array ans = [num1, num2].
.If there are multiple pairs satisfying these conditions,
return the one with the smallest num1 value.
.If no such numbers exist, return [-1, -1]*/

package problems

import "math"

// Sieve of Eratosthenes to find all primes up to `right`(algorithm)
func sieve(limit int) []bool {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

func closestPrimes(left int, right int) []int {
	isPrime := sieve(right)
	primes := []int{}

	// Collect prime numbers in range [left, right]
	for i := left; i <= right; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	// If there are fewer than 2 primes, return [-1, -1]
	if len(primes) < 2 {
		return []int{-1, -1}
	}

	// Find the closest prime pair
	minDiff := math.MaxInt32
	closestPair := []int{-1, -1}
	for i := 1; i < len(primes); i++ {
		diff := primes[i] - primes[i-1]
		if diff < minDiff {
			minDiff = diff
			closestPair = []int{primes[i-1], primes[i]}
		}
	}

	return closestPair

}
