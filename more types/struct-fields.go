package main

import "fmt"

type Vertex3 struct {
	X int
	Y int
}

func main18() {
	v := Vertex3{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
