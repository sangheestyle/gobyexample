package main

import "fmt"

func main() {
	n := map[string]int{"foo": 1, "bar": 2}
	a, b := n["k2"]
	fmt.Println(a, b)

	a, b = n["foo"]
	fmt.Println(a, b)
}
