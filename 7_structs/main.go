package main

import "fmt"

type Foo struct { //public
	MyPublicProperty  string
	myPrivateProperty string
}

func (f Foo) privateMethod() {

}
func (f Foo) PublicMethod() {
	fmt.Println(f.myPrivateProperty)
}
func (f Foo) FailUpdate(bar string) {
	f.myPrivateProperty = bar
}
func (f *Foo) Update(bar string) {
	f.myPrivateProperty = bar
}

type bar struct { //private
	MyPublicProperty  string
	myPrivateProperty string
}

func MakeBar(public, private string) bar {
	return bar{
		myPrivateProperty: private,
		MyPublicProperty:  public,
	}
}

func NewBar(public, private string) *bar {
	return &bar{
		myPrivateProperty: private,
		MyPublicProperty:  public,
	}
}

func main() {
	fmt.Printf("%+v\n", MakeBar("public", "private")) //nice print
	fmt.Printf("%v\n", MakeBar("public", "private"))  //normal print

	//when it makes sens, the functions that are exposed from the package are meant to complement the package name
	// ex package user
	// you would have a constructor that is named New() not NewUser(), as when you would call it it will look like user.New()
	// other constructors should include full name ex NewAddress() which will be called user.NewAddress()
}
