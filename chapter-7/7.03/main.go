package main

import (
	"fmt"
	"strings"
)

type record struct {
	key       string
	valueType string
	data      interface{} // can be "any" as well
}

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

func newRecord(key string, i interface{}) record {
	r := record{}
	r.key = key
	// using a type switch. If i matches any of the types in the cases, then that case will
	// run. Otherwise, default case will run
	switch val := i.(type) {
	case int:
		r.valueType = "int"
		r.data = val
	case bool:
		r.valueType = "bool"
		r.data = val
	case string:
		r.valueType = "string"
		r.data = val
	case person:
		r.valueType = "person"
		r.data = val
	default:
		r.valueType = "unknown"
		r.data = val
	}
	return r
}

func main() {
	// map where key is string and the value can be of any type
	m := make(map[string]interface{})
	a := animal{name: "oreo", category: "cat"}
	p := person{lastName: "Doe", isMarried: false, age: 19}
	m["person"] = p
	m["animal"] = a
	m["age"] = 54
	m["isMarried"] = true
	m["lastName"] = "Smith"

	records := []record{}
	// iterating through all keys in the map and generate records with it
	for key, val := range m {
		r := newRecord(key, val)
		records = append(records, r)
	}

	for _, val := range records {
		fmt.Println("Key: ", val.key)
		fmt.Println("Data: ", val.data)
		fmt.Println("Type: ", val.valueType)
		fmt.Println()
	}

	// more type assertions
	/*
		Stating variable "str" is an empty interface so it can take
		in any value type. Assigning a string value to it but we will
		not be able to use any string methods or operations associated with
		the string type because Go does not know what type it is. We need
		to assert the type and convert the value to string instead of "interface"
		before we can do any string operations.
	*/
	var str interface{} = "the book club"
	v, isValid := str.(string)
	if isValid {
		fmt.Printf("%v is a valid string\n", v)
	}

	/*
		Below is a dangerous type assertion as if it's the unexpected type, it will break your program.
		str2 is of type empty interface and we assign it an int value. However we try to type assert into
		a string value so we can run string operations on it. Without checking if the assertion was valid or not,
		we will break our program. It will also not be caught during compile time as Go relinquishes the responsibility
		of it to the dev.
		See error: "panic: interface conversion: interface {} is int, not string"
	*/
	var str2 interface{} = 49
	v2 := str2.(string)
	fmt.Println(strings.Title(v2))
}
