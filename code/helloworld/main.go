package main

import "fmt"

func main() {
	fmt.Print("Hello World fmt.Print()")
	fmt.Println("Hello World")
	fmt.Printf("Hello World fmt.Printf()")

	fmt.Print(1)
	fmt.Println(2)
	// fmt.Printf(3) -> fmt.Printf must not a integer
	fmt.Printf("3")

	// NB : use double quotes not a single quotes
}
