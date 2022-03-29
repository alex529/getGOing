package main

import "fmt"

type MyEnum int

const (
	Enum0 MyEnum = iota //if no iota int need to be assigned
	Enum1
	Enum2
	Enum3
	Enum4
)

func (m MyEnum) IsValid() bool { //handy function
	switch m {
	case Enum0, Enum1, Enum2, Enum3, Enum4:
		return true
	}
	return false
}

func (m MyEnum) String() string {
	return [...]string{"Enum0", "Enum1", "Enum2", "Enum3", "Enum4"}[m]
}

func main() {
	fmt.Printf("string rep: %s\n", Enum0) //<- implements Stringer
	fmt.Printf("int rep: %d\n", Enum0)

	works := MyEnum(2)
	fmt.Printf("string rep: %s\n", works)

	doesntWorks := MyEnum(5)
	fmt.Printf("string rep: %s\n", doesntWorks) //but it compiles
}
