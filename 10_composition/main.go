package main

import "fmt"

type Foo struct {
	foo string
}

func (f Foo) Hello() {
	fmt.Printf("Hi, I'm foo%s\n", f.foo)
}

type Bar struct {
	bar string
	Foo
}

func main() {
	bar := Bar{
		bar: "bar",
		Foo: Foo{
			foo: ", but also bar",
		},
	}
	bar.Hello()
	//check what happens if bar has also Foo method
}
