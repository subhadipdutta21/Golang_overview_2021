package main

import (
	"fmt"
	"strconv"
)

func getBio(name string, age int, city string) {
	fmt.Println("hi i am " + name + " age " + strconv.Itoa(age) + " from  " + city)

}

func arrayPlayground(data []string) {
	data2 := data
	data2 = append(data2, "testdata")
	fmt.Println(data2)
	fmt.Println(data)
}

func iterator(data []string) {
	for _, item := range data {
		fmt.Println(item)
	}
}

func decissionMaker(a int, b int) {
	if a > b {
		fmt.Println("a is greater")
	} else {
		fmt.Println("b is greater")
	}
}

func main() {
	getBio("subh", 27, "Kolkata")
	testArr := []string{"aaa", "bbb", "ccc"}
	arrayPlayground(testArr)
	iterator(testArr)
	decissionMaker(12, 14)
}
