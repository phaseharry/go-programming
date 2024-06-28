package main

import "fmt"

/*
Go implictly implements interfaces implicity.
See java: class Dog implements Pet (where Pet is an interface).
In Go, any type implements an interface as long as it implements all of the
methods that matches the method set within the defined interface.

Also, a type can satisfy multiple interfaces at the same time. See below, "Speaker"
implements both the "Stringer" and "Speaker" interfaces.

In Go, for interfaces with only one method, typically the name of the
interface would have the `er` suffix. Since there's only the "Speak" method,
it is a Go best practice to have the type be called "Speaker".

Any type that wants to implement te "Speaker" interface, they need to satisfy the method set
(i.e, a type must have defined all of methods within the method set to pass the interface)
*/
type Speaker interface {
	Speak() string
}

type person struct {
	name      string
	age       int
	isMarried bool
}

/*
This method allows the person struct type to pass the "Stringer" interface
and implements it. fmt.Println() takes the Stringer interface
*/
func (p person) String() string {
	return fmt.Sprintf("%v (%v years old). \nMarried status: %v ", p.name, p.age, p.isMarried)
}

/*
This method allows the person struct type to pass the "Speaker" interface we created above
and implements it.
*/
func (p person) Speak() string {
	return "Hi my name is: " + p.name
}

func main() {
	p := person{name: "Cailyn", age: 44, isMarried: false}
	fmt.Println(p.Speak())
	fmt.Println(p)
}
