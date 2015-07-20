package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area1() int {
	return r.width * r.height
}

func (r rect) area2() int {
	return r.width * r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area1())
	fmt.Println("area: ", r.area2())
}
