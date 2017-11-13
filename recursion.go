package main

import "fmt"

func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

func main() {
	defer fmt.Println(factorial(5))

	func() {
		fmt.Println("Anonymous self-executing function")
	}()
}
