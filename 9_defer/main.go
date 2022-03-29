package main

import "fmt"

// similar to the new using
func main() {
	tmp := func(x string) { fmt.Println(x) }

	defer tmp("defer 1")
	defer tmp("defer 2")

	// usually used to cleanup like IDisposable
	// f, _ := os.Open("/tmp/dat")
	// defer f.Close()
}
