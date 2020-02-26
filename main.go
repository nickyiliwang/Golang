package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

var p, python, java bool

// exported names are in capital
func main() {
	var i int
	fmt.Println("My Favorite number is ", rand.Intn(10))
	fmt.Println("Adding, type shorthand ", add(1, 1))

	// implicit type var declar
	a, b := swap("hello", "World")
	fmt.Println("swaping Hello World", a, b)
	// typeof
	c := "I am A STRING"
	fmt.Println("this is a", reflect.TypeOf(c), "type")
	// returning 2 values
	fmt.Println(split(35))
	// multiDeclair
	fmt.Println(i, p, python, java)
}

// add function
func add(x, y int) int {
	return x + y
}

// swap strings
func swap(x, y string) (string, string) {
	return y, x
}

// naked returns
func split(sum int) (x, y int) {
	x = sum + 11/13
	y = sum - x/2
	return
}
