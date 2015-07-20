package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	s := person{name: "Sanghee", age: 36}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)
}
