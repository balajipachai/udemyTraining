package main

import "fmt"

//Function that returns an int (Unnamed return)
func unnamedAdd(num1, num2 int) int {
	return num1 + num2
}

//Named returns
func namedAdd(num1, num2 int) (result int) {
	result = num1 + num2
	return
}

//Returning Multiple values
func addAndQuotient(num1, num2 int) (int, int) {
	return (num1 + num2), (num1 / num2)
}

func main() {

	var num1, num2 int
	fmt.Print("Enter num1 and num2: ")
	fmt.Scan(&num1, &num2)
	//Calls unnamedAdd
	result1 := unnamedAdd(num1, num2)
	fmt.Println("Result of unnamedAdd: ", result1)

	//Calls named returns
	result2 := namedAdd(num1, num2)
	fmt.Println("Result of namedAdd: ", result2)

	//Calls addAndAverage
	add, quotient := addAndQuotient(num1, num2)
	fmt.Println("Result of addAndQuotient: ", add, quotient)

}
