package main

import "fmt"

//Function that takes an int and returns quotient and true
// if the number is even
//else quotient and false
func half(number int) (int, bool) {
	if number%2 == 0 {
		return (number / 2), true
	}
	return (number / 2), false
}

//Variadic function that finds the greatest from a list
func findGreatest(numList ...int) int {
	//Using range traverse the list that has been passed
	//and find the greatest among them
	var greatest, value int
	for _, value = range numList {
		if value >= greatest {
			greatest = value
		}
	}
	return greatest
}

//One more variadic function foo()
func foo(inputList ...int) {
	var value int
	//Traverses the inputList using range
	for _, value = range inputList {
		fmt.Println("In foo: ", value)
	}
}
func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	//Calls the half
	quotient, isSuccess := half(num)
	fmt.Println(quotient, isSuccess)

	//Calling the half using a func expression
	funcCallTohalf := half
	fmt.Println("Using func expression: ")
	fmt.Println(funcCallTohalf(num))

	//Calls greatest
	largestNum := findGreatest(12, 14, 15, 17, 105, 10, 20, 30, 40, 50, 60)
	fmt.Println("Largest of the list: ", largestNum)

	//Value of the expression
	fmt.Println((true && false) || (false && true) || !(false && false))

	//A function foo that can be called in various ways as below
	fmt.Println("First call to foo()")
	foo(1, 2)
	fmt.Println("Second call to foo()")
	foo(1, 2, 3)

	aSlice := []int{1, 2, 3, 4}
	fmt.Println("Third call to foo()")
	foo(aSlice...)
	fmt.Println("Fourth call to foo()")
	foo()
}
