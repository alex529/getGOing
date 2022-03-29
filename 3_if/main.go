package main

import "fmt"

func main() {
	var x string  // string x;
	y := "string" // var y = "string"

	if y != x {
		fmt.Println("not equal") // fmt.Println = Console.WriteLine
	}

	if y != x {
		fmt.Println("not equal")
	} else { //has to be like this
		fmt.Println("equal")
	}

	if areEqual := y == x; !areEqual {
		fmt.Println("not equal")
	}
}
