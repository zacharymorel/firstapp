package main

import (
	"fmt"
	"strconv"
)

var (
	actorName    string = "John Smith"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// SHADOWING
	fmt.Printf("%v, %T\n", actorName, actorName)
	var actorName string = "Jared Smith"
	fmt.Printf("%v, %T\n", actorName, actorName)

	// Conversion
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i) // We need to use string conversion package here because a string is an stream of bites
	fmt.Printf("%v, %T\n", k, k)

}
