package main

import "fmt"

//Referring gobyexample.com/slices
func goByExample() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[2:]
	fmt.Println("sl3: ", l)

	t := []string{"x", "y", "z"}
	fmt.Println("dcl: ", t)
	fmt.Println("****************************************************")
	//Maps
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[125] = "BalaDon"
	fmt.Println(myMap)
	delete(myMap, 125)
	fmt.Println(myMap)

	//Map
	myMap = make(map[int]string)
	myMap[0] = "One"

	fmt.Println(myMap)

	//if value exists in Map
	if val, exists := myMap[0]; exists {
		fmt.Println("Exists and deleted")
		fmt.Println(myMap)
		fmt.Println("val ", val)
		fmt.Println("exists ", exists)
		delete(myMap, 0)
		fmt.Println(myMap)

	} else {
		fmt.Println("Value doesn't exist")
		fmt.Println("val ", val)
		fmt.Println("exists ", exists)
	}

	fmt.Println("****************************************************")

	//create a program that:
	//creates a map using shorthand notation
	//adds an entry to the map
	//changes an entry in the map
	//deletes an entry in the map
	//prints all of the entries in the map using range
	//prints the len of the map
	//uses the comma ok idiom
	shortHandMap := make(map[int]string)
	//adds an entry to the map
	shortHandMap[0] = "Zero"
	shortHandMap[1] = "One"
	shortHandMap[2] = "Two"
	shortHandMap[3] = "Three"
	shortHandMap[4] = "Four"
	shortHandMap[5] = "Five"
	fmt.Println("map1 ", shortHandMap)

	//changes an entry in the map
	shortHandMap[0] = "ChangedZero"
	fmt.Println("map2 ", shortHandMap)

	//deletes an entry in the map
	delete(shortHandMap, 0)
	fmt.Println("map3 ", shortHandMap)

	//prints all of the entries in the map using range
	for _, value := range shortHandMap {
		fmt.Println(value)
	}

	//prints the len of the map
	fmt.Println("Length of Map: ", len(shortHandMap))

	//uses the comma ok idiom
	if val, exists := shortHandMap[4]; exists {
		fmt.Println("Val exists")
		fmt.Println(shortHandMap)
		fmt.Println("val ", val)
		fmt.Println("exists ", exists)
		delete(shortHandMap, 4)
		fmt.Println(shortHandMap)
	} else {
		fmt.Println("Val doesn't exists")
		fmt.Println(shortHandMap)
		fmt.Println("val ", val)
		fmt.Println("exists ", exists)
		delete(shortHandMap, 4)
		fmt.Println(shortHandMap)
	}

	fmt.Println("****************************************************")

	//Structure in GO
	type person struct {
		name string
		age  int
	}

	fmt.Println(person{"Balaji", 25})
	p1 := person{name: "Avinash", age: 24}
	fmt.Printf("Type of person: %T\n", p1)
	fmt.Println("****************************************************")
}

//Slices in go
//Slices can be created in 3 ways, using var, shorthand and make

func main() {
	//Slice created using var
	var varSlice []string

	//Slice created using shorthand
	shorthandSlice := []int{}

	//Slice created using make
	makeSlice := make([]int, 0, 5)

	//Prints all the Slices
	fmt.Println("****************************************************")
	fmt.Println("Before any allocations: ")
	fmt.Println("varSlice: ", varSlice)
	fmt.Println("shorthandSlice: ", shorthandSlice)
	fmt.Println("makeSlice: ", makeSlice)

	fmt.Println("****************************************************")
	//Prints the length and capactiy of all Slices
	fmt.Println("Length of varSlice: ", len(varSlice), " Capacity: ", cap(varSlice))
	fmt.Println("Length of shorthandSlice: ", len(shorthandSlice), " Capacity: ", cap(shorthandSlice))
	fmt.Println("Length of makeSlice: ", len(makeSlice), " Capacity: ", cap(makeSlice))
	fmt.Println("****************************************************")

	fmt.Println("After allocations: ")
	//Adding elements to all Slices
	varSlice = append(varSlice, "Hello Balaji")
	fmt.Println("Length of varSlice: ", len(varSlice), " Capacity: ", cap(varSlice))
	fmt.Println("varSlice: ", varSlice)
	fmt.Println("****************************************************")

	shorthandSlice = append(shorthandSlice, 120, 150, 125, 255, 12587, 147, 1478)
	fmt.Println("Length of shorthandSlice: ", len(shorthandSlice), " Capacity: ", cap(shorthandSlice))
	fmt.Println("shorthandSlice: ", shorthandSlice)
	fmt.Println("****************************************************")

	makeSlice = append(makeSlice, 169, 120, 150, 125, 255, 12548)
	fmt.Println("Length of makeSlice: ", len(makeSlice), " Capacity: ", cap(makeSlice))
	fmt.Println("makeSlice: ", makeSlice)
	fmt.Println("****************************************************")

	goByExample()

}
