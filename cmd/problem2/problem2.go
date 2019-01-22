package main

import "fmt"

func fibonacci(limit int) []int {
	number1 := 1
	number2 := 2
	fibonacci := append([]int{}, number1, number2)

	for number2 < limit {
		buffer := number2
		number2 = number1 + number2
		number1 = buffer

		fibonacci = append(fibonacci, number2)
	}

	return fibonacci
}

func main() {
	fibonacci := fibonacci(4000000)
	sum := 0

	for _, value := range fibonacci {
		if value%2 == 0 {
			sum += value
		}
	}

	fmt.Println("sum:", sum)
}
