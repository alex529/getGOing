package main

import "fmt"

func main() {
	fmt.Println("____________________arrays")
	array1 := [2]string{"1", "2"}
	array2 := [...]string{"1", "2", "2"}
	fmt.Println(array1)
	fmt.Println(array2)

	fmt.Println("____________________slices append")
	slice1 := make([]string, 0)
	fmt.Println(slice1)
	slice1 = append(slice1, "new string") // append
	fmt.Println(slice1)

	fmt.Println("____________________slices delete")
	slice2 := []string{"1", "2", "3", "4"}
	fmt.Println(slice2)
	slice2 = append(slice2[:1], slice2[2:]...) //delete
	fmt.Println(slice2)

	fmt.Println("____________________maps")
	Map := make(map[string]string) //no generics
	Map["exist"] = "tmp"
	tmp := Map["exist"]
	fmt.Println(tmp)

	tmp1 := Map["doesn't exist"]
	fmt.Println(tmp1)
	tmp2, ok := Map["doesn't exist"]
	fmt.Printf("tmp2: '%s', ok: '%t'\n", tmp2, ok)
	if tmp3, ok := Map["exist"]; ok {
		fmt.Println(tmp3)
	}

	delete(Map, "exist")
	_, isPresent := Map["exist"]
	fmt.Println(isPresent)
}
