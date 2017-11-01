package main

import "fmt"

//The final parameter in a function signature may have a type prefixed with ...
//A function with such a parameter is called variadic and may be invoked with zero or
//more arguments for that parameter

func average(sf ...float64) float64 {
	var total float64
	var value float64
	var key int
	//Prints the key value pair of passed list
	for key, value = range sf {
		total += value
		fmt.Printf("sf[%d] = %f\n", key, value)
	}
	return total / float64(len(sf))
}

func main() {
	n := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Average: ", n)
}
