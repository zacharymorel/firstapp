package main

import "fmt"

var (
	actorName    string = "John Smith"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// SHADOWING
	fmt.Printf("%v, %T", actorName, actorName)
	fmt.Printf("\n")
	var actorName string = "Jared Smith"
	fmt.Printf("%v, %T", actorName, actorName)
}
