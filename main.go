package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// exported names are in capital
func main() {
	fmt.Println("My Favorite number is ", rand.Intn(10))
	fmt.Println("Adding, type shorthand ", add(1, 1))

	// implicit type var declar
	a, b := swap("hello", "World")
	fmt.Println("swaping Hello World", a, b)
	// typeof
	c := "I am A STRING"
	fmt.Println("this is a", reflect.TypeOf(c), "type")
}

// add function
func add(x, y int) int {
	return x + y
}

// swap strings
func swap(x, y string) (string, string) {
	return y, x
}
