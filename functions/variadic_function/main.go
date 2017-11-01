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
	//Note that we are calling average with variable number of arguments.
	avg1 := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	avg2 := average(10, 20, 30, 45, 25, 78, 96, 100)
	//average is also called without an argument still it's executed.
	avg3 := average()
	fmt.Println("Average1: ", avg1)
	fmt.Println("Average2: ", avg2)
	fmt.Println("Average3: ", avg3)
}
