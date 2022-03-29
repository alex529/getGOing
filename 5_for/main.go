package main

import "fmt"

func main() {
	array2 := [...]string{"1", "2", "2", "2"}
	for i := 0; i < len(array2); i++ {
		fmt.Println(i, array2[i])
	}
	fmt.Println()

	for k := range array2 {
		fmt.Println(k)
	}
	fmt.Println()

	for k, v := range array2 {
		fmt.Println(k, v)
	}
	fmt.Println()

	for _, v := range array2 {
		fmt.Println(v)
	}
	fmt.Println()

	i := len(array2)
	for i > 0 {
		i--
		fmt.Println(i, array2[i])
	}
}
