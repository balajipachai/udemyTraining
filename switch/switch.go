package main

import (
	"fmt"
)

//swith depending on type
func SwitchOnType(value interface{}) {

	switch value.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float")
	default:
		fmt.Println("unknown type")

	}
}

//Program that shows the usage of switch in go
func main() {
	//Calls SwitchOnType
	SwitchOnType("")
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
		fallthrough
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	default:
		fmt.Println("Number range 1...9 can be printed")

	}

	//Multiple conditions in case
	var bestFriend string
	fmt.Print("Enter your best friend's name: ")
	fmt.Scan(&bestFriend)

	switch bestFriend {

	case "sachu":
		fmt.Println("Hey there Sachu")

		// it checks whether your best friends name is bala or ketan
	case "bala", "ketan":
		fmt.Println("Hey there Sachu or Ketan")

	case "avinash":
		fmt.Println("Hey there Avinash")

	default:
		fmt.Println("Seems your Best Friend is not in the list")

	}

	// switch with no expression
	companyName := "Harbinger"
	switch {
	case len(companyName) == 9:
		fmt.Println("Congratulations !!! you have been selected in harbinger. Length of Harbinger is: ", len(companyName))
	case companyName == "Syntel":
		fmt.Println("Your company is Syntel")
	case companyName == "Google":
		fmt.Println("Your company is Google")
	default:
		fmt.Println("Your company is not in the list")
	}

}

//fallthrough will execute all the cases that follow it.
//if we put fallthrough after case 5 then the values that will be printed are: Five, Six
//fallthrough makes the next case also to be evaluated as true
