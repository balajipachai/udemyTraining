package main

import "fmt"

func main() {
	//Prints 394 in binary using
	fmt.Printf("394 in binary is: %b\n", 394)

	//Prints the type of variable
	a := 125
	fmt.Printf("The type of a is: %T\n", a)

	//Creates avariable of type int initialized to it's zero value.
	//Accept input from user and store it in the above created variable.
	var zero_value int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&zero_value)

	//Loop that creates a lot of UTF-8 characters
	for index := 50; index <= 140; index++ {
		fmt.Println(index, " - ", string(index), " - ", []byte(string(index)))

	}

	//Creates a variable of type string and assigns value to it using backtick
	var backtick_string_variable = `Raw string`
	fmt.Println(backtick_string_variable)

	//Converts a string to a byte of slice and then print it
	var myName = "Balaji"
	fmt.Println(myName, " in bytes is: ", []byte(string(myName)))

	//Prints the length of string
	fmt.Println("Length of myName is: ", len(myName))

}
