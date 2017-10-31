package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
	PB = 1 << (iota * 10)
)

func zero(a *int) {
	*a = 0
	fmt.Println("Address of C in zero: ", &a)

}
func main() {

	fmt.Println("Binary\tDecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
	fmt.Printf("%b\t", PB)
	fmt.Printf("%d\n", PB)

	var l, b int
	fmt.Print("Enter length of rectangle:")
	fmt.Scan(&l)
	fmt.Print("Enter breadth of rectangle:")
	fmt.Scan(&b)

	area := l * b
	fmt.Println("Area: ", area)

	var c int = 125
	fmt.Println(&c)
	var d *int = &c
	fmt.Println(d)
	fmt.Println(*d)
	*d = 100
	fmt.Println(&c)
	fmt.Println(d)
	fmt.Println(c)

	zero(&c)
	fmt.Println("Zeroed C: ", c)
	fmt.Println("Address of C in main: ", &c)

	for index := 0; index < 100; index++ {
		if index%2 == 1 {
			fmt.Printf("%d is Odd\n", index)
		} else {
			fmt.Printf("%d is Even\n", index)
		}
	}

	//Prints UTF-8 characters
	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
