package main

import "fmt"

func sum(numbers ...int) int {
	fmt.Println("numbers looks like an []int: ", numbers)
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("result is:", sum(1, 2, 3, 4, 5))
}

// functionality that is not core related to the struct should be specified as a func not as a method
// but should still be kept in the same file
// like static functions
