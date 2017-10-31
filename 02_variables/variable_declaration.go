package main

import "fmt"

func main() {
	//Shorthand variable declaration
	integer := 10
	bulean := true
	String := "balaji"
	float := 3.14

	fmt.Printf("%v", integer)
	fmt.Println()

	fmt.Printf("%v", bulean)
	fmt.Println()
	fmt.Printf("%v", String)
	fmt.Println()
	fmt.Printf("%v", float)
	fmt.Println()

	//Pritning the type of variables
	fmt.Printf("%T", integer)
	fmt.Println()
	fmt.Printf("%T", bulean)
	fmt.Println()
	fmt.Printf("%T", String)
	fmt.Println()
	fmt.Printf("%T\n", float)

}
