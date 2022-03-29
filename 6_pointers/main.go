package main

import "fmt"

func main() {
	x := 1

	var y *int
	y = &x
	//equivalent to y:= &x

	*y = 2
	fmt.Println(x)
}
