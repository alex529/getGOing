package main

import "fmt"

// Duck typing "If it looks like a duck, and it quacks like a duck, then it is a duck"

type Speaker interface { //Stringer
	Say(string)
}

type Person struct {
	name string
}

func (p *Person) Say(message string) {
	fmt.Printf("%s: %s\n", p.name, message)
}

func joke(s Speaker) {
	s.Say("Knock, Knock")
}

func main() {
	john := &Person{
		name: "John",
	}
	joke(john)

	// jim := Person{
	// 	name: "John",
	// }
	// joke(jim) //doesn't work make it
}

//interface tend to be very small, you can compose them io package:
// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }
// type ReadWriter interface {
// 	Reader
// 	Writer
// }
