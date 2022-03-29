package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string `json:"FULL_NAME"` //if `json:"<field_name>"` not supplied it uses the name of the propriety
	Age      int    `json:"AGE,omitempty"`
	Nickname string `json:"-"` //there can be custom attributes here they are used with reflection
}

func main() {
	bob := Person{Name: "Bob", Age: 0, Nickname: "Bil"}
	fmt.Printf("%+v\n", bob)

	p, _ := json.Marshal(bob)
	fmt.Println(string(p))

	var who Person
	_ = json.Unmarshal(p, &who) //due to the fact that generics wasn't a thing you see a lot of this kind of pattern where the variable needs to be passed so you get the type with the help of reflection
	fmt.Printf("%+v", who)
}

// when unmarshalling large strings you can use json.RawMessage type to stop unmarshalling nested objects
