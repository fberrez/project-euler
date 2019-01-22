package main

import (
	"fmt"
)

func isPalindrome(N int) bool {
	reverse := 0
	buffer := N

	for buffer != 0 {
		reverse = reverse * 10
		reverse = reverse + buffer%10
		buffer = buffer / 10
	}

	return (N == reverse)
}

func palindrome() (int, int, int) {
	max := 0
	maxFactor1 := 0
	maxFactor2 := 0

	for factor1 := 100; factor1 < 1000; factor1++ {
		for factor2 := 100; factor2 < 1000; factor2++ {
			product := factor1 * factor2

			if product > max && isPalindrome(product) {
				max = product
				maxFactor1 = factor1
				maxFactor2 = factor2
			}
		}
	}

	return max, maxFactor1, maxFactor2
}

func main() {
	max, factor1, factor2 := palindrome()
	fmt.Println(fmt.Sprintf("largest palindrome made from the product of two 3-digit (%d*%d) numbers is %d", factor1, factor2, max))
}
