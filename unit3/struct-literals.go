package main

import "fmt"

type Vertex4 struct {
	X, Y int
}

var (
	v1 = Vertex4{1, 2}  // has type Vertex4
	v2 = Vertex4{X: 1}  // Y:0 is implicit
	v3 = Vertex4{}      // X:0 and Y:0
	p  = &Vertex4{1, 2} // has type *Vertex4
)

func main19() {
	fmt.Println(v1, p, v2, v3)
}
