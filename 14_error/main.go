package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Person struct {
	Name     string `json:"FULL_NAME"` //if `json:"<field_name>"` not supplied it uses the name of the propriety
	Age      int    `json:"AGE,omitempty"`
	Nickname string `json:"-"` //there can be custom attributes here they are used with reflection
}

// 2 types
// panic("msg") -> recover() seldom used similar to exceptions
// error interface

var ErrJSON = errors.New("it's a json error") // similar to declaring own exception type

type ComplexError struct {
	detail string
}

func (ce ComplexError) Error() string { // implement error interface
	return ce.detail
}

func run() error {
	p, err := json.Marshal(Person{Name: "Bob", Age: 0, Nickname: "Bil"})
	if err != nil {
		return err
	}

	var bob Person
	if err := json.Unmarshal(p, &bob); err != nil {
		return err
	}
	// if err := json.Unmarshal([]byte("lots of errors"), &bob); err != nil {
	// 	return ErrJSON
	// }

	// if err := json.Unmarshal([]byte("lots of complex errors"), &bob); err != nil {
	// 	return ComplexError{
	// 		detail: "a lot of details",
	// 	}
	// }

	// //shadowing example
	// if err := json.Unmarshal([]byte("lots of errors"), &bob); err != nil {
	// 	err = fmt.Errorf("error while unmarshalling bob with msg %s", "custom msg here")
	// } //please fix

	return err
}

func main() {
	err := run()
	if errors.Is(err, ErrJSON) {
		fmt.Println("it was a JSON error")
	}
	var complexErr ComplexError
	if errors.As(err, &complexErr) {
		fmt.Printf("this was a complexError with the following details: %s", complexErr.detail)
	}

	if err == nil {
		fmt.Println("You're an amazing programmer, no errors were returned")
	}
}
