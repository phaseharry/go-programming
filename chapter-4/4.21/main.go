package main

import (
	"errors"
	"fmt"
)

/*
An interface defines a value that has to have the functions to associated with it to be a valid type to be passed into a function.
If a value gets passed in does not have the functions on the interface defined then Go will scream at you at compile time.
Having interface{} as the type means a value of any type can be passed in. This is the same as "any" type. However, you can't use the values
type in processing unless assert it. Type asserting it allows you to use the "any" or "interface{}" underlying value's type to do processing.
It removes the responsibility from Go's compiler to type check it and the developer needs to handle the potential failures during runtime
*/
func doubler(v interface{}) (string, error) {
	// checking if value can be type asserted to int, if it's an int, then process it
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}
	// checking if value can be type asserted to string, if it's an string, then process it
	if s, ok := v.(string); ok {
		return s + s, nil
	}

	return "", errors.New("unsupported type passed")
}

func main() {
	res, _ := doubler(5)
	fmt.Println("5: ", res)
	res, _ = doubler("yum")
	fmt.Println("yum: ", res)
	_, err := doubler(true)
	fmt.Println("true: ", err)
}
