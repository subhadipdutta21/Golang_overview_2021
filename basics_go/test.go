package main

import (
	"fmt"
	"strconv"
)

func getBio(name string, city string, pin int) string { //fun call with args passing
	return "Hi i'm " + name + " City " + city + " and pin" + strconv.Itoa(pin)
}

func getArray() []string {
	fruits := []string{"apple", "orange"} //array declaration
	return fruits
}

// structure / object example
type Person struct {
	Name    string
	Job     string
	Address Address
}
type Address struct {
	City  string
	State string
}

func main() {
	x := 5
	y := 3

	// iterating example on array
	users := []string{"subh", "anant", "malhar", "satya", "sajid"}
	test := []int{}

	for _, item := range users {
		fmt.Println(item)
	}

	for i := 0; i <= 5; i++ {
		test = append(test, i)
	}

	fmt.Println("populated array", test)

	fmt.Println("-------------------iterating on key val pair------------------------------")
	// iterating on map (key val pairs)
	items := map[string]string{"subh": "sd@gmail.com", "anant": "anant@gmail.com", "manab": "manab@gmail.com"}

	for k, v := range items {
		fmt.Println(k + " : " + v)
	}

	fmt.Println("-------------------------------------------------")
	if x > y {
		fmt.Println("x is gt y")
	} else {
		fmt.Println("y is gt x")
	}

	// structure / object example
	p := Person{
		Name: "Steve",
		Job:  "Software dev",
		Address: Address{
			City:  "Gotham",
			State: "NY",
		},
	}

	fmt.Println("_strut example__", p)
	fmt.Println("_func with args passing example__", getBio("subh", "Kol", 560024)) //fun call with args passing
	fmt.Println("_array example__", getArray())
}
