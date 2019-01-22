package main

import "fmt"

func findMultiples(limit int) []int {
	multiples := []int{}

	for i := 0; i < limit; i++ {
		if i%3 == 0 {
			multiples = append(multiples, i)
		} else if i%5 == 0 {
			multiples = append(multiples, i)
		}
	}

	return multiples
}
func main() {
	multiples := findMultiples(1000)
	sum := 0

	for _, multiple := range multiples {
		sum += multiple
	}

	fmt.Println("sum:", sum)
}
