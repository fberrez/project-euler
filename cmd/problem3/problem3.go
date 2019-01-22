package main

import (
	"fmt"
	"math"
)

func findPrimeFactors(N int) []int {
	primeFactors := []int{}
	p := 2

	for N >= int(math.Pow(float64(p), 2)) {
		if N%p == 0 {
			primeFactors = append(primeFactors, p)
			N = N / p
		} else {
			p++
		}
	}

	primeFactors = append(primeFactors, N)

	return primeFactors
}

func main() {
	N := 600851475143
	primeFactors := findPrimeFactors(N)
	largestPrimeFactor := primeFactors[len(primeFactors)-1]

	fmt.Println(fmt.Sprintf("largest prime factor of %d: %v", N, largestPrimeFactor))
}
