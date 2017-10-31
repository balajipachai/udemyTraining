package main

import "fmt"

var x int = 100

const q = 100

const temp string = "hi bala"

func main() {

	fmt.Println("Hello World using Atom")
	foo()
	fmt.Printf("42 in binary %b\n", 42)
	fmt.Println(x)
	fmt.Println(y)

	q := 200
	fmt.Println("Changed value of q: ", q)
	temp := "hey there"
	fmt.Println(temp)
	print()
}

func foo() {
	fmt.Println(x)

}

func print() {
	fmt.Println(temp)
	fmt.Println(q)
	const (
		A = iota
		B
		C
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	const (
		D = iota
		E
		F
	)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}

var y = 125
