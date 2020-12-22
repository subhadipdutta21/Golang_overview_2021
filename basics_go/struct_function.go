package main

import (
	"fmt"
)

type Student struct {
	name   string
	grades []int
	age    int
}

// getter
func (s Student) getAge() int {
	return s.age
}

// setter
func (s *Student) setAge(age int) {
	s.age = age
}

func main() {
	s1 := Student{"Subh", []int{20, 25, 45, 35, 23}, 19}
	s2 := Student{"Subh", []int{22, 45, 56, 25, 66}, 25}
	fmt.Println(s1, s2)

	fmt.Println(s1.getAge(), s2.getAge())

	s1.setAge(45)
	fmt.Println(s1)

}
